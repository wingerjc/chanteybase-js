export function GetQueryParam(name) {
  return new URLSearchParams(document.location.search).get(name);
}
