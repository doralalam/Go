# Arrays
->  Array is a data structure that holds values from similar data objects, unlike struct which stores data from different data objects




# Slices
->  Slices are like window to an array where we can access part of the array through slicing.
->  Slices were like pointers to an array because, when we create a slice out of an array, instead of copying the data, slice directs to the memory location of that array. So, any changes we made to slices will be applied directly on the array.
-> len() is a built-in function in GO that gives the number of elements stored in an array
-> cap() is yet another built-in function in GO that gives the max. no of elements in the array (neglecting the ones that are ommitted on its left and considers all the values on its right side)


# Maps
-> Map is a data structure used to group the data