// An example of using maps within maps. This example demonstrates a few
// features. Firstly, creating gBPF map specifications in pure Go
// (typically you'd see them being generated from a loaded ELF).
// Additionally, creating maps and placing them in other maps (with
// dynamically sized inner maps).
package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/khulnasoft/gbpf"
	"github.com/khulnasoft/gbpf/rlimit"
)

const BPF_F_INNER_MAP = 0x1000

func main() {
	// Allow the current process to lock memory for gBPF resources.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal(err)
	}

	// We're creating a map spec in pure Go here, but a map spec like
	// this can be loaded from an ELF too.
	outerMapSpec := gbpf.MapSpec{
		Name:       "outer_map",
		Type:       gbpf.ArrayOfMaps,
		KeySize:    4, // 4 bytes for u32
		ValueSize:  4,
		MaxEntries: 5, // We'll have 5 maps inside this map
		Contents:   make([]gbpf.MapKV, 5),
		InnerMap: &gbpf.MapSpec{
			Name:      "inner_map",
			Type:      gbpf.Array,
			KeySize:   4, // 4 bytes for u32
			ValueSize: 4, // 4 bytes for u32

			// This flag is required for dynamically sized inner maps.
			// Added in linux 5.10.
			Flags: BPF_F_INNER_MAP,

			// We set this to 1 now, but this inner map spec gets copied
			// and altered later.
			MaxEntries: 1,
		},
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// For each entry we want to create in the outer map...
	for i := uint32(0); i < outerMapSpec.MaxEntries; i++ {
		// Copy the inner map spec
		innerMapSpec := outerMapSpec.InnerMap.Copy()

		// Randomly generate inner map length
		innerMapSpec.MaxEntries = uint32(r.Intn(50) + 1) // Can't be zero.

		// populate the inner map contents
		innerMapSpec.Contents = make([]gbpf.MapKV, innerMapSpec.MaxEntries)

		for j := range innerMapSpec.Contents {
			innerMapSpec.Contents[uint32(j)] = gbpf.MapKV{Key: uint32(j), Value: uint32(0xCAFE)}
		}

		// Create the inner map
		innerMap, err := gbpf.NewMap(innerMapSpec)
		if err != nil {
			log.Fatalf("inner_map: %v", err)
		}
		// In this example we close all references to maps before exit.
		// But typically you may actually want to hold on to the map
		// reference so that you control the lifecycle of the map. For
		// the inner (nested) map though, it's safe to close the file
		// descriptor in userspace once the outer map holds a reference
		// in the kernel.
		defer innerMap.Close()

		// Inner map is created successfully and lives in the kernel,
		// let's add it to the contents of the outer map spec.
		outerMapSpec.Contents[i] = gbpf.MapKV{Key: i, Value: innerMap}
	}

	// All inner maps are created and inserted into the outer map spec,
	// time to create the outer map.
	outerMap, err := gbpf.NewMap(&outerMapSpec)
	if err != nil {
		log.Fatalf("outer_map: %v", err)
	}
	defer outerMap.Close()

	// The outer map is created successfully and lives happily in the
	// kernel. Let's iterate over the map in the kernel to see what's
	// been made.
	mapIter := outerMap.Iterate()
	var outerMapKey uint32
	var innerMapID gbpf.MapID
	for mapIter.Next(&outerMapKey, &innerMapID) {
		// With maps that contain maps, performing a lookup doesn't give
		// you the map directly, instead it gives you an ID, which you
		// can then use to get a full map pointer.
		innerMap, err := gbpf.NewMapFromID(innerMapID)
		if err != nil {
			log.Fatal(err)
		}

		innerMapInfo, err := innerMap.Info()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("outerMapKey %d, innerMap.Info: %+v", outerMapKey, innerMapInfo)
	}
}
