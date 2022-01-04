const { readFileSync, writeFileSync } = require("fs");

const packageJSONFilePath = "sdk/nodejs/bin/package.json"

const packageJSON = JSON.parse(
    readFileSync(packageJSONFilePath, { encoding: "utf-8" })
);

if (!packageJSON.scripts) {
    packageJSON.scripts = {};
}
var name = packageJSON.name;
if (name.lastIndexOf("/") !== -1) {
    name = name.substring(name.lastIndexOf("/") + 1);
}
packageJSON.scripts[
    "install"
] = `pulumi plugin install resource ${name} ${packageJSON.version} --server ${packageJSON.pulumi.pluginDownloadURL}`;

writeFileSync(
    packageJSONFilePath,
    JSON.stringify(packageJSON, null, 4)
);
