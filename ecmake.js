var sh = std.LoadPlugin('../sh/sh');

function Test(args) {
    console.log(JSON.stringify(sh.RunV("go", "test", "-v", "./...", "-cover")))
}

SetTargets({
    "Test": Test,
    "TestWithRace": function (args) {
        sh.RunV("go", "test", "-race", "-v", "./...", "-cover")
    },
    "Release": function (args) {
        console.log(sh.RunV("bash", "-c", "curl -sL https://git.io/goreleaser | bash"));
    }
});

