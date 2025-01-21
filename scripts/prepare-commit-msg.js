const fs = require('fs');
const commitMsgFile = process.argv[2];

const commitMsg = fs.readFileSync(commitMsgFile, 'utf-8');
fs.writeFileSync(commitMsgFile, `${commitMsg.trim()} (modified by git hook)`);