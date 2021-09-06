export function CreateQueryString(attribMap) {
  var result = "?";
  console.log(attribMap);
  for (const [key, value] of Object.entries(attribMap)) {
    if (result.length > 1) {
      result += "&";
    }
    result += encodeURIComponent(key) + "=" + encodeURIComponent(value);
  }
  return result;
}
