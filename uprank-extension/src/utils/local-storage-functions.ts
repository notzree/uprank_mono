export function setWithExpiry(key, value: any, ttl) {
  const now = new Date()
  // `item` is an object which contains the original data as well as the expiry time
  const item = {
    value: value,
    expiry: now.getTime() + ttl
  }
  localStorage.setItem(key, JSON.stringify(item))
}

export function getWithExpiry(key) {
  const itemStr = localStorage.getItem(key)
  // If the key does not exist, return null
  if (!itemStr) {
    return null
  }
  const item = JSON.parse(itemStr)
  const now = new Date()
  // Compare the expiry time of the item with the current time
  if (now.getTime() > item.expiry) {
    // If the item is expired, delete the item from storage and return null
    localStorage.removeItem(key)
    return null
  }
  return item.value
}
