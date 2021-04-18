var sh = require('sh');
var http = require('http');

function Test(args) {
    console.log(JSON.stringify(sh.RunV("go", "test", "-v", "./...", "-cover")));

    var resp = http.Get('https://google.com')
    var resp = http.Post('https://ptsv2.com/t/dgl3x-1618752767/post', "application/json", {"test": "fish"})
    console.log(JSON.stringify(resp, null, 2));
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

