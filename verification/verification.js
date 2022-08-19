const starkwareCrypto = require('@starkware-industries/starkware-crypto-utils');
const fs = require('fs');
const BN = require('bn.js');

// 10000 signatures
let filePath = "./signatures_produced_by_current_lib.txt"
// produced by this lib
let signatures = loadSignatures()

function loadSignatures(){
    let objStr = ""
    try {
        objStr = fs.readFileSync(`${filePath}`, 'utf8')
    } catch (err) {
        console.error(`failed to read file, ${err}`)
        return []
    }

    let lines = objStr.split(/[\n]/g)
    let temSig = {}
    let rlt = []
    for (let index = 0; index < lines.length; index+= 7) {
        temSig["hash"] = lines[index]
        temSig["privateKey"] = lines[index+1]
        temSig["xPublicKey"] = lines[index+2]
        temSig["yPublicKey"] = lines[index+3]
        temSig["r"] = lines[index+4]
        temSig["s"] = lines[index+5]

        rlt.push(temSig)
    }

    return rlt
}

function verifyAllSignatures(){
    let allPass = true
    let failedCount = 0

    for (let index = 0; index < signatures.length; index++) {
        let sig = signatures[index];

        let privKeyBN = new BN(sig.privateKey, 10)
        let privKeyHex = privKeyBN.toString(16)
        let hashBN = new BN(sig.hash, 10)
        let hashHex = hashBN.toString(16)
        let rBN = new BN(sig.r, 10)
        let sBN = new BN(sig.s, 10)

        let keyPair = starkwareCrypto.ec.keyFromPrivate(privKeyHex, 'hex')
        let publicKey = starkwareCrypto.ec.keyFromPublic(keyPair.getPublic(true, 'hex'), 'hex');

        let verifyRlt = starkwareCrypto.verify(publicKey, hashHex, {r: rBN, s: sBN})
        if(!verifyRlt){
            allPass = false
            failedCount++
            console.log(`failed to verify the sig ${JSON.stringify(sig)}`)
        }{
            console.log(`success to verify the ${index}-th signature`)
        }
    }

    if(allPass){
        console.log(`success to verify all ${signatures.length} signatures.`)
    }else{
        console.log(`signature verification pass: ${failedCount} / ${signatures.length}`)
    }
}

verifyAllSignatures()


