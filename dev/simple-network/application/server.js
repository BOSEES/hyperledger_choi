// 웹서버 모듈 포함 
const express = require('express');
const app = express();
const bodyParser = require('body-parser');
//웹 서버 설정

const { FileSystemWallet, Gateway } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const ccpPath = path.resolve(__dirname, '..', 'network' ,'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

//express 설정
const port = 3333;
const host = "0.0.0.0";

app.use(express.static(path.join(__dirname,"views")));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended:false}));


// 페이지 라우팅
app.get("/", async(req, res) => res.send("홈페이지"));
app.post("/mate", async(req, res) => {
    //인자 가져오기
    const name = req.body.name;
    //인증서 가져오기
    const walletPath = path.join(process.cwd(), 'wallet');
    const wallet = new FileSystemWallet(walletPath);
    console.log(`Wallet path: ${walletPath}`);
    
    const userExists = await wallet.exists('user1');
    if (!userExists) {
        console.log('An identity for the user "user1" does not exist in the wallet');
        console.log('Run the registerUser.js application before retrying');
        return;
    }
    // 체인코드 수행
    const gateway = new Gateway();
    await gateway.connnect(ccp, {wallet,identity:"user1", discovery: {enabled : false}});
    const network = await gateway.getNetwork("mychannel");
    const contract = network.getContract("teamate");
    await contract.submitTransaction("addUser", name);
    console.log("transaction has been submitted");
    await gateway.disconnect();
    
    res.status(200).send("transaction has been submitted");
    //
});
// app.post("/mate", (req, res) => res.send("mate"));
app.post("/score", (req, res) => res.send("score"));


// 2.1 mate 추가 arg: name ,mate post
// 2.2 mate 조회 arg: name ,mate get
// 2.2 mate project 추가 arg:name, project, score /score post
// 서버시작

app.listen(port, host, () => {
    console.log(`Server listen :${host}:${port}`);
});