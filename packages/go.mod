module example.com/bankPackage

go 1.23.6


// to create a module
// go mod init <path>
// this can be later used as prefix while importing the packages
// example import "example.com/bankPackage/fileOperatioins"
// NOTE: Different packages must be in different directories. Otherwise, go will face the ambiguity