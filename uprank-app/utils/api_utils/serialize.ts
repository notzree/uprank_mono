//Serializes dates to ISO strings in an object and returns the object

export function serializeDates<T extends Record<string, any>>(obj: T): T {
    for (const key in obj) {
        const value = obj[key];
        if (value as object instanceof Date){
            obj[key] = obj[key].toISOString();
        }
        else if (value as object instanceof Object){
            obj[key] = serializeDates(value);
        }
        
    }
    return obj;
}