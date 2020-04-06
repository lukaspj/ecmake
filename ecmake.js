function Test(args) {
    console.log("Tester");
    try {
        sh.Output("go");
    } catch (e) {
        console.log(JSON.stringify(e));
        console.log(e.toString());
    }
    console.log(sh.Output("go", "version"));
}

SetTargets({
    "Test": Test
});

