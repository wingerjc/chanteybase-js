load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "b34cbe1a7514f5f5487c3bfee7340a4496713ddf4f119f7a225583d6cafd793a",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.21.1/rules_go-v0.21.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.21.1/rules_go-v0.21.1.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "86c6d481b3f7aedc1d60c1c211c6f76da282ae197c3b3160f54bd3a8f847896f",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.1/bazel-gazelle-v0.19.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.1/bazel-gazelle-v0.19.1.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_jmoiron_sqlx",
    commit = "2ba0fc60eb4a54030f3a6d73ff0a047349c7eeca",
    importpath = "github.com/jmoiron/sqlx",
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    commit = "9bdaffc12bf8be15afceb51bb60851edd4afdff5",
    importpath = "github.com/mattn/go-sqlite3",
)

go_repository(
    name = "com_github_lib_pq",
    commit = "4604d39ddc9f62e2ca152114c73aec99b77a3468",
    importpath = "github.com/lib/pq",
)

http_archive(
    name = "io_bazel_rules_rust",
    sha256 = "b5d4d1c7609714dfef821355f40353c58aa1afb3803401b3442ed2355db9b0c7",
    strip_prefix = "rules_rust-8d2b4eeeff9dce24f5cbb36018f2d60ecd676639",
    urls = [
        # Master branch as of 2020-11-10
        "https://github.com/bazelbuild/rules_rust/archive/8d2b4eeeff9dce24f5cbb36018f2d60ecd676639.tar.gz",
    ],
)

load("@io_bazel_rules_rust//rust:repositories.bzl", "rust_repositories")

rust_repositories(version = "1.49.0", edition = "2018")

load("@io_bazel_rules_rust//:workspace.bzl", "rust_workspace")

rust_workspace()

load("//cargo:crates.bzl", "raze_fetch_remote_crates")

raze_fetch_remote_crates()