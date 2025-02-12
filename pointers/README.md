# Pointers is one of the important features in Go
->  Pointers stores the address of the variable instead of the actual value
# Advantages:
1.  In go, whenever a value is passed through a variable to a function, go by default copies the value.
    With complex project, this will create a significant impact in the memory.
    If we use pointer instead, functions will be passed with address of the variable, where any changes made to the variable address will directly applied on the pointed varibale instead of making copies
2. With pointers, only one value is stored in the memory and it is passed through
3. While using pointers in the function, no return values are required
4. Pointers lead to less code, due to the elimination of copy nature

# Disadvantages:
1. Sometimes, we unknowingly overwrites the data using dereferences