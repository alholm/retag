// Package retag provides ability to change tags of structures' fields in runtime
// without copying of the data. It may be helpful in next cases:
//
//  - Automatic tags generation;
//  - Different views of the one data;
//  - Fixing of leaky abstractions with minimal boilerplate code
//    when application has layers of abstractions and model is
//    separated from storages and presentation layers.
//
// This package is still experimental and subject to change.
//
package retag
