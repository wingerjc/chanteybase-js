import m from "mithril";
import { GetQueryParam } from "./util";

export default function PersonPage(initialVnode) {
  var data = {};
  var id = GetQueryParam("id");
  m.request({
    method: "GET",
    url: "../api/person/:id",
    params: { id: id },
  })
    .then(function (result) {
      data.result = result;
      data.fetched = true;
      console.log(data);
    })
    .catch(function (result) {
      data.error = result;
      data.fetched = true;
      console.log(data);
    });
  return {
    view: function (vnode) {
      if (data.error) {
        return m("h1", data.error.response.message);
      } else if (data.result) {
        const d = data.result[0];
        return m("", [
          m("p", "id: " + d["id"]),
          m("p", "First Name: " + d["first-name"]),
          m("p", "Last Name: " + d["last-name"]),
          m("br"),
        ]);
      }
    },
  };
}
