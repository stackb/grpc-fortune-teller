workspace(name = "com_github_staackb_fortune_teller")

#####################################################################
# RULES_GO
#####################################################################

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "161c91485b007c6bf51c0e81808cf4ee2ded299d",
)

http_archive(
    name = "com_github_scele_rules_go_dep",
    urls = ["https://github.com/scele/rules_go_dep/archive/49a5e4ca9f6a16c9b4c930a51ce3a537498bb4e1.tar.gz"],
    strip_prefix = "rules_go_dep-49a5e4ca9f6a16c9b4c930a51ce3a537498bb4e1",
    sha256 = "f170d3d6f55e216f1493f975cde6c489d7070da2a8a41fd4de9812d96f4fb38b",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
load("@com_github_scele_rules_go_dep//dep:dep.bzl", "dep_import")

go_register_toolchains(go_version = "1.9")

go_rules_dependencies()

dep_import(
    name = "godeps",
    prefix = "github.com/stackb/fortune-teller",
    gopkg_lock = "//:Gopkg.lock",
)

load("@godeps//:Gopkg.bzl", "go_deps")

go_deps()

#############################################################
# RULES_DOCKER
#############################################################

RULES_DOCKER_VERSION = "553d5506bb7325185950f91533b967da8f5bc536"

http_archive(
    name = "io_bazel_rules_docker",
    url = "https://github.com/bazelbuild/rules_docker/archive/%s.zip" % RULES_DOCKER_VERSION,
    strip_prefix = "rules_docker-" + RULES_DOCKER_VERSION,
    sha256 = "e0b3d966f2a5c0fe921b6294df7c823afa63b4c439f0a7f3b9da3ed6534bab83",
)

load(
    "@io_bazel_rules_docker//container:container.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    go_image_repositories = "repositories",
)

go_image_repositories()
