import m from "mithril";
import header from "../components/header";
import footer from "../components/footer";

export default function (contentComponent) {
  console.log(contentComponent);
  return {
    view: function (initialVnode) {
      return m("div", { class: "layout" }, [
        m(header),
        m("div", "navbar"),
        m(contentComponent),
        m(footer),
      ]);
    },
  };
}
