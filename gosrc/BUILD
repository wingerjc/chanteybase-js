load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/wingerjc/chanteybase-js/go
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/wingerjc/chanteybase-js/go",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_jmoiron_sqlx//:go_default_library",
        "@com_github_mattn_go_sqlite3//:go_default_library",
    ],
)

go_binary(
    name = "go",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
