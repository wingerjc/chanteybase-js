import m from "mithril";

export default {
  view: function () {
    return m("div", { class: "header" }, [
      m("h1", { class: "header__main-text" }, "ChanteyBase"),
    ]);
  },
};
