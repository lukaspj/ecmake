const io = require('io');
const sh = require('sh');
const http = require('http');

function Test(args) {
    console.log(JSON.stringify(sh.RunV("go", "test", "-v", "./...", "-cover")));

    var resp = http.Get('https://google.com')
    var resp = http.Post('https://ptsv2.com/t/dgl3x-1618752767/post', "application/json", {"test": "fish"})
    console.log(JSON.stringify(resp, null, 2));

    io.Walk("modules", (path, fileinfo, error) => {
        console.log(path + " -> " + fileinfo.Name() + ", " + fileinfo.IsDir());
    })
}

SetTargets({
    "Test": Test,
    "TestWithRace": (args) => {
        sh.RunV("go", "test", "-race", "-v", "./...", "-cover")
    },
    "Release": (args) => {
        console.log(sh.RunV("bash", "-c", "curl -sL https://git.io/goreleaser | bash"));
    }
});

