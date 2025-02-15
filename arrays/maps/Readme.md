# Maps

-> Map is a data structure, also used to group the data

# Maps vs Structs
# Advantages of maps over struct as both use same key value pair functionality
->  In maps, we can use any type of data as keys such as int, float, string, array, struct etc.,
->  We can add or delete new key value pair in maps whereas, in structs, once defined, we can't add or delete
    Example:
    type websites struct{
        google string
        aws string
        linkedin string
    }
    In the above code, we can't add a new key value pair.
    Struct is only useful, whenever we had a clear idea about the requirement
    Otherwise, maps are useful for custom labeled data.