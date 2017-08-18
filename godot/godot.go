// package name: libgodot
package godot

/*
#cgo CXXFLAGS: -I/usr/local/include -std=c11
#cgo LDFLAGS: -Wl,-unresolved-symbols=ignore-all
#include <stddef.h>
#include <stdlib.h>
#include <godot_nativescript.h>

// Type definitions for any function pointers. Some of these are not defined in
// the godot headers when they are part of a typedef struct.
typedef void (*create_func)(godot_object *, void *);
typedef void (*destroy_func)(godot_object *, void *, void *);
typedef void (*free_func)(void *);
typedef godot_variant (*method)(godot_object *, void *, void *, int, godot_variant **);

// Forward declarations of gateway functions defined in cgateway.go.
void *go_create_func_cgo(godot_object *, void *); // Forward declaration.
void *go_destroy_func_cgo(godot_object *, void *); // Forward declaration.
void *go_free_func_cgo(void *); // Forward declaration.
godot_variant go_method_func_cgo(godot_object *obj, void *method_data, void *user_data, int num_args, godot_variant **args);
*/
import "C"

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"unsafe"
)

const (
	tagPrefix        = "godot"
	defaultBaseClass = "Node"
)

// Exposed is a base structure for any structure that will be exposed to Godot.
// You should embed this structure in your own custom structs.
type Exposed struct {
	GDParent string `godot:"_inherits"`
}

// ClassConstructor is any function that will build and return a class to be registered
// with Godot.
type ClassConstructor func() interface{}

// function signature for Godot classes
type gdMethod func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer, numArgs C.int, args **C.godot_variant)

type Method func(godotObject *Object, methodData, userData interface{}, numArgs int, args []*Variant)

/** Library entry point **/
// godot_gdnative_init is the library entry point. When the library is loaded
// this method will be called by Godot.
//export godot_gdnative_init
func godot_gdnative_init(options *C.godot_gdnative_init_options) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(Log)
	log.Println("Initializing Go library.")

	// Translate the C struct into a Go struct.
	goOptions := &GodotGDNativeInitOptions{
		InEditor:      bool(options.in_editor),
		CoreAPIHash:   uint64(options.core_api_hash),
		EditorAPIHash: uint64(options.editor_api_hash),
		NoAPIHash:     uint64(options.no_api_hash),
	}

	// Call user defined init if it was set.
	if godotGDNativeInit != nil {
		godotGDNativeInit(goOptions)
	}
}

/** Library de-initialization **/
// godot_gdnative_terminate is the library's de-initialization method. When
// Godot unloads the library, this method will be called.
//export godot_gdnative_terminate
func godot_gdnative_terminate(options *C.godot_gdnative_terminate_options) {
	log.Println("De-initializing Go library.")
}

/** Script entry (Registering all the classes and stuff) **/
// godot_nativescript_init is the script's entrypoint. It is called by Godot
// when a script is loaded. It is responsible for registering all the classes
// and stuff. The `unsafe.Pointer` type is used to represent a null C pointer.
//export godot_nativescript_init
func godot_nativescript_init(desc unsafe.Pointer) {
	log.Println("Initializing script")

	// Loop through our registered classes and register them with the Godot API.
	for _, constructor := range godotConstructorsToRegister {
		// Use the constructor to build a class to inspect the given structure.
		class := constructor()

		// Get the type of the given struct, and get its name as a string and C string.
		classType := reflect.TypeOf(class)
		classString := strings.Replace(classType.String(), "*main.", "", 1) // TODO: Just remove * instead.
		classCString := C.CString(classString)
		log.Println("Registering class:", classString)

		// Add the type to our internal type registry. This is used so the constructor
		// function can create the correct kind of struct.
		constructorRegistry[classString] = constructor
		typeRegistry[classString] = classType

		// Look at the struct tags for "_inherits" to get the base class.
		baseClass := defaultBaseClass
		classValuePtr := reflect.ValueOf(class)
		classValue := classValuePtr.Elem()
		log.Println("  Found", classValue.NumField(), "struct fields.")
		for i := 0; i < classValue.NumField(); i++ {
			// Get the struct field.
			field := classValue.Field(i)

			// Ignore fields that don't have the same type as a string
			if field.Type() != reflect.TypeOf("") {
				continue
			}

			// Check if the struct field tag is "inherits"
			fieldTag, _ := classValue.Type().Field(i).Tag.Lookup(tagPrefix)
			if fieldTag != "_inherits" {
				continue
			}

			baseClass = field.Interface().(string)
			log.Println("  Found inheritance in struct:", baseClass)
		}
		baseClassCString := C.CString(baseClass)
		log.Println("  Using Base Class:", baseClass)

		// TODO: Look at the struct tags for methods to register.
		// TODO: OR look for methods that follow the GodotMethod type.

		// Set up our create function C struct
		var createFunc C.godot_instance_create_func
		var destroyFunc C.godot_instance_destroy_func

		// *** CREATE FUNC ***
		// Use a pointer to our gateway function.
		// GDCALLINGCONV void *(*create_func)(godot_object *, void *);
		createFunc.create_func = (C.create_func)(unsafe.Pointer(C.go_create_func_cgo))
		// void *method_data;
		createFunc.method_data = unsafe.Pointer(classCString)
		// GDCALLINGCONV void (*free_func)(void *);
		createFunc.free_func = (C.free_func)(unsafe.Pointer(C.go_free_func_cgo))

		// *** DESTROY FUNC ***
		// GDCALLINGCONV void (*destroy_func)(godot_object *, void *, void *);
		destroyFunc.destroy_func = (C.destroy_func)(unsafe.Pointer(C.go_destroy_func_cgo))
		// void *method_data;
		destroyFunc.method_data = unsafe.Pointer(classCString)
		// GDCALLINGCONV void (*free_func)(void *);
		destroyFunc.free_func = (C.free_func)(unsafe.Pointer(C.go_free_func_cgo))

		// Register our class.
		C.godot_nativescript_register_class(desc, classCString, baseClassCString, createFunc, destroyFunc)

		// TODO: Register the methods of the class.
		// Loop through our class's methods that are attached to it.
		log.Println("  Looking at methods:")
		log.Println("    Found", classType.NumMethod(), "methods")
		for i := 0; i < classType.NumMethod(); i++ {
			classMethod := classType.Method(i)
			log.Println("  Found method:", classMethod.Name)

			// Check to see if the method's name is a special Godot method.
			methodString := classMethod.Name
			if methodName, ok := methodMap[classMethod.Name]; ok {
				log.Println("    Mapped method to:", methodName)
				methodString = methodName
			}

			// Set up registering a method
			var method C.godot_instance_method

			// *** METHOD STRUCTURE ***
			// GDCALLINGCONV godot_variant (*method)(godot_object *, void *, void *, int, godot_variant **);
			method.method = (C.method)(unsafe.Pointer(C.go_method_func_cgo))
			// void *method_data;
			method.method_data = unsafe.Pointer(classCString)
			// GDCALLINGCONV void (*free_func)(void *);
			method.free_func = (C.free_func)(unsafe.Pointer(C.go_free_func_cgo))

			// Set up the method attributes.
			var attr C.godot_method_attributes
			attr.rpc_type = C.GODOT_METHOD_RPC_MODE_DISABLED

			// Set up the method name
			methodCString := C.CString(methodString)

			// Register a method.
			C.godot_nativescript_register_method(desc, classCString, methodCString, attr, method)
		}

	}
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in cgateway.go.
//export go_create_func
func go_create_func(godotObject *C.godot_object, methodData unsafe.Pointer) unsafe.Pointer {
	// Cast our pointer to a *char, so we can cast it to a Go string.
	str := (*C.char)(methodData)
	className := C.GoString(str)
	log.Println("Create function called for:", className)

	// Look up our class constructor by its class name in the registry.
	constructor := constructorRegistry[className]

	// Create a new instance of the object.
	class := constructor()
	instanceString := fmt.Sprintf("%p", &class)
	instanceCString := C.CString(instanceString)
	log.Println("Created new object instance:", class, "with instance address:", instanceString)

	// Add the instance to our instance registry.
	instanceRegistry[instanceString] = class

	// Return the instance string. This will be passed to the method function as userData, so we
	// can look up the instance in our registry. Normally you would pass a pointer to the instance
	// itself, but we should never pass Go structures to C, as the Go garbage collector might
	// reap the object prematurely.
	return unsafe.Pointer(instanceCString)
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in cgateway.go.
//export go_destroy_func
func go_destroy_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer) {
	log.Println("Destroy function called!")

	// Look up the instance based on the userData string.
	instanceCString := (*C.char)(userData)
	instanceString := C.GoString(instanceCString)

	// Remove it from our instanceRegistry so it can be garbage collected.
	log.Println("Destroying instance:", instanceString)
	delete(instanceRegistry, instanceString)
}

//export go_free_func
func go_free_func(methodData unsafe.Pointer) {
	log.Println("Free function called!")
}

//godot_variant go_method_func_cgo(godot_object *obj, void *method_data, void *user_data, int num_args, godot_variant **args)
//export go_method_func
func go_method_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer, numArgs C.int, args **C.godot_variant) C.godot_variant {
	log.Println("Method was called!")
	log.Println("  godotObject:", godotObject)
	log.Println("  methodData:", methodData)
	log.Println("  userData:", userData)
	log.Println("  numArgs:", int(numArgs))
	log.Println("  args:", args)

	// Get the object instance based on the instance string given in userData.
	instanceCString := (*C.char)(userData)
	instanceString := C.GoString(instanceCString)
	class := instanceRegistry[instanceString]
	log.Println("  instance:", class)
	log.Println("  instanceString:", instanceString)

	// Create a slice of godot_variant arguments
	argsSlice := []*C.godot_variant{}

	// Get the size of each godot_variant object pointer.
	size := unsafe.Sizeof(*args)

	// Panic if something's wrong.
	if int(numArgs) > 500 {
		panic("Wtf, 500+ arguments?")
	}

	// If we have arguments, append the first argument.
	if int(numArgs) > 0 {
		argsSlice = append(argsSlice, *args)

		// Loop through all our arguments.
		for i := 0; i < int(numArgs); i++ {
			// Convert the pointer into a uintptr so we can perform artithmetic on it.
			arrayPtr := uintptr(unsafe.Pointer(args))

			// Add the size of the godot_variant pointer to our array pointer to get the position
			// of the next argument.
			arg := (*C.godot_variant)(unsafe.Pointer(arrayPtr + size))

			// Check to see what type this variant is
			variantType := C.godot_variant_get_type(arg)
			log.Println("Argument variant type:", variantType)

			// Append the argument to our slice.
			argsSlice = append(argsSlice, arg)
		}
	}

	var ret C.godot_variant
	C.godot_variant_new_nil(&ret)

	return ret
}
