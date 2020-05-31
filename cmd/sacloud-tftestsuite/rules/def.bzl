def _sacloud_tf_config_test_impl(ctx):
    kicker = ctx.actions.declare_file("%s_kicker.sh" % ctx.label.name)
    ctx.actions.expand_template(
        template = ctx.file._wrapper_template,
        output = kicker,
        substitutions = {
            "{executable_binary}": ctx.executable._test_suite.short_path,
            "{config_file}": ','.join([src.short_path for src in ctx.files.srcs]),
        },
        is_executable = True,
    )
    runfiles = ctx.runfiles(collect_data = True, files = [kicker, ctx.executable._test_suite] + ctx.files.srcs)
    return [DefaultInfo(executable = kicker, runfiles = runfiles)]
sacloud_tf_config_test = rule(
    implementation = _sacloud_tf_config_test_impl,
    test = True,
    attrs = {
#        "src": attr.label(allow_single_file = True),
        "srcs": attr.label_list(allow_files = [".tf"]),
        "_test_suite": attr.label(
            default = Label("//cmd/sacloud-tftestsuite"),
            executable = True,
            cfg = "target",
        ),
        "_wrapper_template": attr.label(
            allow_single_file = True,
            default = "kicker.tpl",
        )
    },
)