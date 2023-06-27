// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const MIPSStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/cannon/MIPS.sol:MIPS\",\"label\":\"oracle\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_contract(IPreimageOracle)1001\"}],\"types\":{\"t_contract(IPreimageOracle)1001\":{\"encoding\":\"inplace\",\"label\":\"contract IPreimageOracle\",\"numberOfBytes\":\"20\"}}}"

var MIPSStorageLayout = new(solc.StorageLayout)

var MIPSDeployedBin = "0x608060405234801561001057600080fd5b50600436106100415760003560e01c8063155633fe146100465780637dc0d1d01461006757806398bb138314610098575b600080fd5b61004e61016c565b6040805163ffffffff9092168252519081900360200190f35b61006f610174565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b61015a600480360360408110156100ae57600080fd5b8101906020810181356401000000008111156100c957600080fd5b8201836020820111156100db57600080fd5b803590602001918460018302840111640100000000831117156100fd57600080fd5b91939092909160208101903564010000000081111561011b57600080fd5b82018360208201111561012d57600080fd5b8035906020019184600183028401116401000000008311171561014f57600080fd5b509092509050610190565b60408051918252519081900360200190f35b634000000081565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600061019a611a62565b608081146101a757600080fd5b604051610600146101b757600080fd5b606486146101c457600080fd5b61016684146101d257600080fd5b6101ef565b8035602084810360031b9190911c8352920192910190565b8560806101fe602082846101d7565b9150915061020e602082846101d7565b9150915061021e600482846101d7565b9150915061022e600482846101d7565b9150915061023e600482846101d7565b9150915061024e600482846101d7565b9150915061025e600482846101d7565b9150915061026e600482846101d7565b9150915061027e600182846101d7565b9150915061028e600182846101d7565b9150915061029e600882846101d7565b6020810190819052909250905060005b60208110156102d0576102c3600483856101d7565b90935091506001016102ae565b505050806101200151156102ee576102e6610710565b915050610708565b6101408101805160010167ffffffffffffffff1690526060810151600090610316908261081e565b9050603f601a82901c16600281148061033557508063ffffffff166003145b15610382576103788163ffffffff1660021461035257601f610355565b60005b60ff16600261036b856303ffffff16601a6108e6565b63ffffffff16901b610959565b9350505050610708565b6101608301516000908190601f601086901c81169190601587901c16602081106103a857fe5b602002015192508063ffffffff851615806103c957508463ffffffff16601c145b156103fa578661016001518263ffffffff16602081106103e557fe5b6020020151925050601f600b86901c166104b1565b60208563ffffffff16101561045d578463ffffffff16600c148061042457508463ffffffff16600d145b8061043557508463ffffffff16600e145b15610446578561ffff169250610458565b6104558661ffff1660106108e6565b92505b6104b1565b60288563ffffffff1610158061047957508463ffffffff166022145b8061048a57508463ffffffff166026145b156104b1578661016001518263ffffffff16602081106104a657fe5b602002015192508190505b60048563ffffffff16101580156104ce575060088563ffffffff16105b806104df57508463ffffffff166001145b156104fe576104f0858784876109c4565b975050505050505050610708565b63ffffffff60006020878316106105635761051e8861ffff1660106108e6565b9095019463fffffffc861661053481600161081e565b915060288863ffffffff161015801561055457508763ffffffff16603014155b1561056157809250600093505b505b600061057189888885610b4d565b63ffffffff9081169150603f8a16908916158015610596575060088163ffffffff1610155b80156105a85750601c8163ffffffff16105b15610687578063ffffffff16600814806105c857508063ffffffff166009145b156105ff576105ed8163ffffffff166008146105e457856105e7565b60005b89610959565b9b505050505050505050505050610708565b8063ffffffff16600a1415610620576105ed858963ffffffff8a1615611213565b8063ffffffff16600b1415610642576105ed858963ffffffff8a161515611213565b8063ffffffff16600c1415610659576105ed6112f8565b60108163ffffffff16101580156106765750601c8163ffffffff16105b15610687576105ed81898988611770565b8863ffffffff1660381480156106a2575063ffffffff861615155b156106d15760018b61016001518763ffffffff16602081106106c057fe5b63ffffffff90921660209290920201525b8363ffffffff1663ffffffff146106ee576106ee84600184611954565b6106fa85836001611213565b9b5050505050505050505050505b949350505050565b6000610728565b602083810382015183520192910190565b60806040518061073a60208285610717565b9150925061074a60208285610717565b9150925061075a60048285610717565b9150925061076a60048285610717565b9150925061077a60048285610717565b9150925061078a60048285610717565b9150925061079a60048285610717565b915092506107aa60048285610717565b915092506107ba60018285610717565b915092506107ca60018285610717565b915092506107da60088285610717565b60209091019350905060005b6020811015610808576107fb60048386610717565b90945091506001016107e6565b506000815281810382a081900390209150505b90565b60008061082a836119f0565b9050600384161561083a57600080fd5b602081019035610857565b60009081526020919091526040902090565b8460051c8160005b601b8110156108af5760208501943583821c60011680156108875760018114610898576108a5565b6108918285610845565b93506108a5565b6108a28483610845565b93505b505060010161085f565b5060805191508181146108ca57630badf00d60005260206000fd5b5050601f8516601c0360031b1c63ffffffff1691505092915050565b600063ffffffff8381167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80850183169190911c821615159160016020869003821681901b830191861691821b92911b0182610943576000610945565b815b90861663ffffffff16179250505092915050565b6000610963611a62565b5060e08051610100805163ffffffff90811690935284831690526080918516156109b357806008018261016001518663ffffffff16602081106109a257fe5b63ffffffff90921660209290920201525b6109bb610710565b95945050505050565b60006109ce611a62565b5060806000600463ffffffff881614806109ee57508663ffffffff166005145b15610a645760008261016001518663ffffffff1660208110610a0c57fe5b602002015190508063ffffffff168563ffffffff16148015610a3457508763ffffffff166004145b80610a5c57508063ffffffff168563ffffffff1614158015610a5c57508763ffffffff166005145b915050610ae1565b8663ffffffff1660061415610a825760008460030b13159050610ae1565b8663ffffffff1660071415610a9f5760008460030b139050610ae1565b8663ffffffff1660011415610ae157601f601087901c1680610ac55760008560030b1291505b8063ffffffff1660011415610adf5760008560030b121591505b505b606082018051608084015163ffffffff169091528115610b27576002610b0c8861ffff1660106108e6565b63ffffffff90811690911b8201600401166080840152610b39565b60808301805160040163ffffffff1690525b610b41610710565b98975050505050505050565b6000603f601a86901c81169086166020821015610f215760088263ffffffff1610158015610b815750600f8263ffffffff16105b15610c28578163ffffffff1660081415610b9d57506020610c23565b8163ffffffff1660091415610bb457506021610c23565b8163ffffffff16600a1415610bcb5750602a610c23565b8163ffffffff16600b1415610be25750602b610c23565b8163ffffffff16600c1415610bf957506024610c23565b8163ffffffff16600d1415610c1057506025610c23565b8163ffffffff16600e1415610c23575060265b600091505b63ffffffff8216610e7157601f600688901c16602063ffffffff83161015610d455760088263ffffffff1610610c6357869350505050610708565b63ffffffff8216610c835763ffffffff86811691161b9250610708915050565b8163ffffffff1660021415610ca75763ffffffff86811691161c9250610708915050565b8163ffffffff1660031415610cd2576103788163ffffffff168763ffffffff16901c826020036108e6565b8163ffffffff1660041415610cf6575050505063ffffffff8216601f84161b610708565b8163ffffffff1660061415610d1a575050505063ffffffff8216601f84161c610708565b8163ffffffff1660071415610d45576103788763ffffffff168763ffffffff16901c886020036108e6565b8163ffffffff1660201480610d6057508163ffffffff166021145b15610d72578587019350505050610708565b8163ffffffff1660221480610d8d57508163ffffffff166023145b15610d9f578587039350505050610708565b8163ffffffff1660241415610dbb578587169350505050610708565b8163ffffffff1660251415610dd7578587179350505050610708565b8163ffffffff1660261415610df3578587189350505050610708565b8163ffffffff1660271415610e0f575050505082821719610708565b8163ffffffff16602a1415610e42578560030b8760030b12610e32576000610e35565b60015b60ff169350505050610708565b8163ffffffff16602b1415610e6b578563ffffffff168763ffffffff1610610e32576000610e35565b50610f1c565b8163ffffffff16600f1415610e945760108563ffffffff16901b92505050610708565b8163ffffffff16601c1415610f1c578063ffffffff1660021415610ebd57505050828202610708565b8063ffffffff1660201480610ed857508063ffffffff166021145b15610f1c578063ffffffff1660201415610ef0579419945b60005b6380000000871615610f12576401fffffffe600197881b169601610ef3565b9250610708915050565b6111ac565b60288263ffffffff16101561108b578163ffffffff1660201415610f6e57610f658660031660080260180363ffffffff168563ffffffff16901c60ff1660086108e6565b92505050610708565b8163ffffffff1660211415610fa457610f658660021660080260100363ffffffff168563ffffffff16901c61ffff1660106108e6565b8163ffffffff1660221415610fd55750505063ffffffff60086003851602811681811b198416918316901b17610708565b8163ffffffff1660231415610fee578392505050610708565b8163ffffffff1660241415611022578560031660080260180363ffffffff168463ffffffff16901c60ff1692505050610708565b8163ffffffff1660251415611057578560021660080260100363ffffffff168463ffffffff16901c61ffff1692505050610708565b8163ffffffff1660261415610f1c5750505063ffffffff60086003851602601803811681811c198416918316901c17610708565b8163ffffffff16602814156110c35750505060ff63ffffffff60086003861602601803811682811b9091188316918416901b17610708565b8163ffffffff16602914156110fc5750505061ffff63ffffffff60086002861602601003811682811b9091188316918416901b17610708565b8163ffffffff16602a141561112d5750505063ffffffff60086003851602811681811c198316918416901c17610708565b8163ffffffff16602b1415611146578492505050610708565b8163ffffffff16602e141561117a5750505063ffffffff60086003851602601803811681811b198316918416901b17610708565b8163ffffffff1660301415611193578392505050610708565b8163ffffffff16603814156111ac578492505050610708565b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f696e76616c696420696e737472756374696f6e00000000000000000000000000604482015290519081900360640190fd5b600061121d611a62565b506080602063ffffffff86161061129557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f76616c6964207265676973746572000000000000000000000000000000000000604482015290519081900360640190fd5b63ffffffff8516158015906112a75750825b156112d557838161016001518663ffffffff16602081106112c457fe5b63ffffffff90921660209290920201525b60808101805163ffffffff808216606085015260049091011690526109bb610710565b6000611302611a62565b506101e051604081015160808083015160a084015160c09094015191936000928392919063ffffffff8616610ffa141561137a5781610fff81161561134c57610fff811661100003015b63ffffffff84166113705760e08801805163ffffffff838201169091529550611374565b8395505b50611723565b8563ffffffff16610fcd14156113965763400000009450611723565b8563ffffffff1661101814156113af5760019450611723565b8563ffffffff1661109614156113e757600161012088015260ff83166101008801526113d9610710565b97505050505050505061081b565b8563ffffffff16610fa314156115a15763ffffffff83166114075761159c565b63ffffffff8316600514156115795760006114298363fffffffc16600161081e565b6000805460208b01516040808d015181517fe03110e1000000000000000000000000000000000000000000000000000000008152600481019390935263ffffffff16602483015280519495509293849373ffffffffffffffffffffffffffffffffffffffff9093169263e03110e19260448082019391829003018186803b1580156114b357600080fd5b505afa1580156114c7573d6000803e3d6000fd5b505050506040513d60408110156114dd57600080fd5b508051602090910151909250905060038516600481900382811015611500578092505b508185101561150d578491505b8260088302610100031c925082600882021b9250600180600883600403021b036001806008858560040301021b0391508119811690508381198616179450505061155f8563fffffffc16600185611954565b60408a018051820163ffffffff169052965061159c915050565b63ffffffff8316600314156115905780945061159c565b63ffffffff9450600993505b611723565b8563ffffffff16610fa414156116755763ffffffff8316600114806115cc575063ffffffff83166002145b806115dd575063ffffffff83166004145b156115ea5780945061159c565b63ffffffff83166006141561159057600061160c8363fffffffc16600161081e565b60208901519091506003841660040383811015611627578093505b83900360089081029290921c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600193850293841b0116911b1760208801526000604088015293508361159c565b8563ffffffff16610fd71415611723578163ffffffff16600314156117175763ffffffff831615806116ad575063ffffffff83166005145b806116be575063ffffffff83166003145b156116cc576000945061159c565b63ffffffff8316600114806116e7575063ffffffff83166002145b806116f8575063ffffffff83166006145b80611709575063ffffffff83166004145b15611590576001945061159c565b63ffffffff9450601693505b6101608701805163ffffffff808816604090920191909152905185821660e09091015260808801805180831660608b01526004019091169052611764610710565b97505050505050505090565b600061177a611a62565b5060806000601063ffffffff88161415611799575060c08101516118f1565b8663ffffffff16601114156117b95763ffffffff861660c08301526118f1565b8663ffffffff16601214156117d3575060a08101516118f1565b8663ffffffff16601314156117f35763ffffffff861660a08301526118f1565b8663ffffffff16601814156118285763ffffffff600387810b9087900b02602081901c821660c08501521660a08301526118f1565b8663ffffffff166019141561185a5763ffffffff86811681871602602081901c821660c08501521660a08301526118f1565b8663ffffffff16601a14156118a5578460030b8660030b8161187857fe5b0763ffffffff1660c0830152600385810b9087900b8161189457fe5b0563ffffffff1660a08301526118f1565b8663ffffffff16601b14156118f1578463ffffffff168663ffffffff16816118c957fe5b0663ffffffff90811660c0840152858116908716816118e457fe5b0463ffffffff1660a08301525b63ffffffff84161561192657808261016001518563ffffffff166020811061191557fe5b63ffffffff90921660209290920201525b60808201805163ffffffff80821660608601526004909101169052611949610710565b979650505050505050565b600061195f836119f0565b9050600384161561196f57600080fd5b6020810190601f8516601c0360031b83811b913563ffffffff90911b1916178460051c60005b601b8110156119e55760208401933582821c60011680156119bd57600181146119ce576119db565b6119c78286610845565b94506119db565b6119d88583610845565b94505b5050600101611995565b505060805250505050565b60ff81166103800261016681019036906104e601811015611a5c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611aed6023913960400191505060405180910390fd5b50919050565b6040805161018081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101919091526101608101611ac8611acd565b905290565b604051806104000160405280602090602082028036833750919291505056fe636865636b207468617420746865726520697320656e6f7567682063616c6c64617461a164736f6c6343000706000a"

var MIPSDeployedSourceMap = "1027:26641:0:-:0;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1479:45;;;:::i;:::-;;;;;;;;;;;;;;;;;;;1867:29;;;:::i;:::-;;;;;;;;;;;;;;;;;;;16655:5668;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;16655:5668:0;;-1:-1:-1;16655:5668:0;-1:-1:-1;16655:5668:0;:::i;:::-;;;;;;;;;;;;;;;;1479:45;1514:10;1479:45;:::o;1867:29::-;;;;;;:::o;16655:5668::-;16733:7;16752:18;;:::i;:::-;16866:4;16859:5;16856:15;16846:2;;16952:1;16949;16942:12;16846:2;17000:4;16994:11;17007;16991:28;16981:2;;17090:1;17087;17080:12;16981:2;17150:3;17132:16;17129:25;17119:2;;17241:1;17238;17231:12;17119:2;17297:3;17283:12;17280:21;17270:2;;17387:1;17384;17377:12;17270:2;17416:416;;;17650:24;;17638:2;17634:13;;;17631:1;17627:21;17623:52;;;;17692:20;;17746:21;;;17800:18;;;17494:338::o;:::-;17854:16;17911:4;17950:18;17965:2;17962:1;17959;17950:18;:::i;:::-;17942:26;;;;18000:18;18015:2;18012:1;18009;18000:18;:::i;:::-;17992:26;;;;18054:17;18069:1;18066;18063;18054:17;:::i;:::-;18046:25;;;;18110:17;18125:1;18122;18119;18110:17;:::i;:::-;18102:25;;;;18154:17;18169:1;18166;18163;18154:17;:::i;:::-;18146:25;;;;18202:17;18217:1;18214;18211;18202:17;:::i;:::-;18194:25;;;;18246:17;18261:1;18258;18255;18246:17;:::i;:::-;18238:25;;;;18290:17;18305:1;18302;18299;18290:17;:::i;:::-;18282:25;;;;18336:17;18351:1;18348;18345;18336:17;:::i;:::-;18328:25;;;;18386:17;18401:1;18398;18395;18386:17;:::i;:::-;18378:25;;;;18434:17;18449:1;18446;18443;18434:17;:::i;:::-;18489:2;18482:10;;18472:21;;;;18426:25;;-1:-1:-1;18482:10:0;-1:-1:-1;18572:1:0;18557:77;18582:2;18579:1;18576:9;18557:77;;;18615:17;18630:1;18627;18624;18615:17;:::i;:::-;18607:25;;-1:-1:-1;18607:25:0;-1:-1:-1;18600:1:0;18593:9;18557:77;;;18561:14;;;18670:5;:12;;;18666:109;;;18751:13;:11;:13::i;:::-;18744:20;;;;;18666:109;18784:10;;;:15;;18798:1;18784:15;;;;;18861:8;;;;-1:-1:-1;;18853:20:0;;-1:-1:-1;18853:7:0;:20::i;:::-;18839:34;-1:-1:-1;18900:10:0;18908:2;18900:10;;;;18969:1;18959:11;;;:26;;;18974:6;:11;;18984:1;18974:11;18959:26;18955:332;;;19212:64;19223:6;:11;;19233:1;19223:11;:20;;19241:2;19223:20;;;19237:1;19223:20;19212:64;;19274:1;19245:25;19248:4;19255:10;19248:17;19267:2;19245;:25::i;:::-;:30;;;;19212:10;:64::i;:::-;19205:71;;;;;;;18955:332;19512:15;;;;19323:9;;;;19452:4;19446:2;19438:10;;;19437:19;;;19512:15;19537:2;19529:10;;;19528:19;19512:36;;;;;;;;;;;;-1:-1:-1;19573:5:0;19592:11;;;;;:29;;;19607:6;:14;;19617:4;19607:14;19592:29;19588:756;;;19676:5;:15;;;19692:5;19676:22;;;;;;;;;;;;;;-1:-1:-1;;19735:4:0;19729:2;19721:10;;;19720:19;19588:756;;;19769:4;19760:6;:13;;;19756:588;;;19878:6;:13;;19888:3;19878:13;:30;;;;19895:6;:13;;19905:3;19895:13;19878:30;:47;;;;19912:6;:13;;19922:3;19912:13;19878:47;19874:229;;;19980:4;19987:6;19980:13;19975:18;;19874:229;;;20067:21;20070:4;20077:6;20070:13;20085:2;20067;:21::i;:::-;20062:26;;19874:229;19756:588;;;20133:4;20123:6;:14;;;;:32;;;;20141:6;:14;;20151:4;20141:14;20123:32;:50;;;;20159:6;:14;;20169:4;20159:14;20123:50;20119:225;;;20235:5;:15;;;20251:5;20235:22;;;;;;;;;;;;;20230:27;;20328:5;20320:13;;20119:225;20369:1;20359:6;:11;;;;:25;;;;;20383:1;20374:6;:10;;;20359:25;20358:42;;;;20389:6;:11;;20399:1;20389:11;20358:42;20354:117;;;20423:37;20436:6;20444:4;20450:5;20457:2;20423:12;:37::i;:::-;20416:44;;;;;;;;;;;20354:117;20500:13;20481:16;20636:4;20626:14;;;;20622:402;;20697:21;20700:4;20707:6;20700:13;20715:2;20697;:21::i;:::-;20691:27;;;;20751:10;20746:15;;20781:16;20746:15;20795:1;20781:7;:16::i;:::-;20775:22;;20825:4;20815:6;:14;;;;:32;;;;;20833:6;:14;;20843:4;20833:14;;20815:32;20811:203;;;20904:4;20892:16;;20998:1;20990:9;;20811:203;20622:402;;21049:10;21062:26;21070:4;21076:2;21080;21084:3;21062:7;:26::i;:::-;21091:10;21062:39;;;;-1:-1:-1;21183:4:0;21176:11;;;21211;;;:24;;;;;21234:1;21226:4;:9;;;;21211:24;:39;;;;;21246:4;21239;:11;;;21211:39;21207:759;;;21270:4;:9;;21278:1;21270:9;:22;;;;21283:4;:9;;21291:1;21283:9;21270:22;21266:132;;;21346:37;21357:4;:9;;21365:1;21357:9;:21;;21373:5;21357:21;;;21369:1;21357:21;21380:2;21346:10;:37::i;:::-;21339:44;;;;;;;;;;;;;;;21266:132;21416:4;:11;;21424:3;21416:11;21412:109;;;21478:28;21487:5;21494:2;21498:7;;;;21478:8;:28::i;21412:109::-;21538:4;:11;;21546:3;21538:11;21534:109;;;21600:28;21609:5;21616:2;21620:7;;;;;21600:8;:28::i;21534:109::-;21705:4;:11;;21713:3;21705:11;21701:72;;;21743:15;:13;:15::i;21701:72::-;21864:4;21856;:12;;;;:27;;;;;21879:4;21872;:11;;;21856:27;21852:104;;;21910:31;21921:4;21927:2;21931;21935:5;21910:10;:31::i;21852:104::-;22018:6;:14;;22028:4;22018:14;:28;;;;-1:-1:-1;22036:10:0;;;;;22018:28;22014:85;;;22087:1;22062:5;:15;;;22078:5;22062:22;;;;;;;;;:26;;;;:22;;;;;;:26;22014:85;22137:9;:26;;22150:13;22137:26;22133:84;;22179:27;22188:9;22199:1;22202:3;22179:8;:27::i;:::-;22290:26;22299:5;22306:3;22311:4;22290:8;:26::i;:::-;22283:33;;;;;;;;;;;;;16655:5668;;;;;;;:::o;2189:1501::-;2230:11;2374:206;;;2474:2;2470:13;;;2460:24;;2454:31;2443:43;;2514:13;;2553;;;2425:155::o;:::-;2605:4;2650;2644:11;2694:5;2724:21;2742:2;2738;2732:4;2724:21;:::i;:::-;2712:33;;;;2781:21;2799:2;2795;2789:4;2781:21;:::i;:::-;2769:33;;;;2842:20;2860:1;2856:2;2850:4;2842:20;:::i;:::-;2830:32;;;;2905:20;2923:1;2919:2;2913:4;2905:20;:::i;:::-;2893:32;;;;2956:20;2974:1;2970:2;2964:4;2956:20;:::i;:::-;2944:32;;;;3011:20;3029:1;3025:2;3019:4;3011:20;:::i;:::-;2999:32;;;;3062:20;3080:1;3076:2;3070:4;3062:20;:::i;:::-;3050:32;;;;3113:20;3131:1;3127:2;3121:4;3113:20;:::i;:::-;3101:32;;;;3166:20;3184:1;3180:2;3174:4;3166:20;:::i;:::-;3154:32;;;;3223:20;3241:1;3237:2;3231:4;3223:20;:::i;:::-;3211:32;;;;3278:20;3296:1;3292:2;3286:4;3278:20;:::i;:::-;3337:2;3327:13;;;;-1:-1:-1;3266:32:0;-1:-1:-1;3391:1:0;3376:84;3401:2;3398:1;3395:9;3376:84;;;3438:20;3456:1;3452:2;3446:4;3438:20;:::i;:::-;3426:32;;-1:-1:-1;3426:32:0;-1:-1:-1;3419:1:0;3412:9;3376:84;;;3380:14;3497:1;3493:2;3486:13;3548:5;3544:2;3540:14;3533:5;3528:27;3639:14;;;3622:32;;;-1:-1:-1;;2189:1501:0;;:::o;13651:1493::-;13722:10;13744:14;13761:23;13773:10;13761:11;:23::i;:::-;13744:40;;13830:1;13824:4;13820:12;13817:2;;;13845:1;13842;13835:12;13817:2;13959;13947:15;;;13904:20;13975:199;;;;14022:12;;;14116:2;14109:13;;;;14157:2;14144:16;;;14004:170::o;:::-;14206:4;14203:1;14199:12;14236:4;14384:1;14369:319;14394:2;14391:1;14388:9;14369:319;;;14509:2;14497:15;;;14450:20;14540:12;;;14554:1;14536:20;14573:42;;;;14637:1;14632:42;;;;14529:145;;14573:42;14590:23;14605:7;14599:4;14590:23;:::i;:::-;14582:31;;14573:42;;14632;14649:23;14667:4;14658:7;14649:23;:::i;:::-;14641:31;;14529:145;-1:-1:-1;;14412:1:0;14405:9;14369:319;;;14373:14;14722:4;14716:11;14701:26;;14797:7;14791:4;14788:17;14778:2;;14878:10;14875:1;14868:21;14916:2;14913:1;14906:13;14778:2;-1:-1:-1;;15050:2:0;15040:13;;15028:10;15024:30;15021:1;15017:38;15079:16;15097:10;15075:33;;-1:-1:-1;;13651:1493:0;;;;:::o;1903:280::-;1962:6;1997:16;;;;2005:7;;;;1997:16;;;;;;1996:23;;;;;2011:1;2054:2;:8;;;2048:15;;;;;2047:21;;2046:30;;;;;;;2102:8;;2101:14;1996:23;2153:21;;2173:1;2153:21;;;2164:6;2153:21;2139:10;;;;;:36;;-1:-1:-1;;;1903:280:0;;;;:::o;12138:472::-;12205:7;12224:18;;:::i;:::-;-1:-1:-1;12323:8:0;;;12352:12;;;12341:23;;;;;;;12374:19;;;;;12284:4;;12407:12;;;12403:171;;12462:6;12471:1;12462:10;12435:5;:15;;;12451:7;12435:24;;;;;;;;;:37;;;;:24;;;;;;:37;12403:171;12590:13;:11;:13::i;:::-;12583:20;12138:472;-1:-1:-1;;;;;12138:472:0:o;9498:1235::-;9591:7;9610:18;;:::i;:::-;-1:-1:-1;9670:4:0;9693:17;9743:1;9733:11;;;;;:26;;;9748:6;:11;;9758:1;9748:11;9733:26;9729:599;;;9798:9;9810:5;:15;;;9826:5;9810:22;;;;;;;;;;;;;9798:34;;9868:2;9862:8;;:2;:8;;;:23;;;;;9874:6;:11;;9884:1;9874:11;9862:23;9861:54;;;;9897:2;9891:8;;:2;:8;;;;:23;;;;;9903:6;:11;;9913:1;9903:11;9891:23;9846:69;;9729:599;;;;9936:6;:11;;9946:1;9936:11;9932:396;;;9991:1;9984:2;9978:14;;;;9963:29;;9932:396;;;10021:6;:11;;10031:1;10021:11;10017:311;;;10075:1;10069:2;10063:13;;;10048:28;;10017:311;;;10105:6;:11;;10115:1;10105:11;10101:227;;;10183:4;10177:2;10169:10;;;10168:19;10206:8;10202:42;;10243:1;10237:2;10231:13;;;10216:28;;10202:42;10270:3;:8;;10277:1;10270:8;10266:43;;;10308:1;10301:2;10295:14;;;;10280:29;;10266:43;10101:227;;10354:8;;;;;10383:12;;;;10372:23;;;;;10437:259;;;;10523:1;10498:21;10501:4;10508:6;10501:13;10516:2;10498;:21::i;:::-;:26;;;;;;;10484:41;;10493:1;10484:41;10469:56;:12;;;:56;10437:259;;;10649:12;;;;;10664:1;10649:16;10634:31;;;;10437:259;10713:13;:11;:13::i;:::-;10706:20;9498:1235;-1:-1:-1;;;;;;;;9498:1235:0:o;22329:5337::-;22416:6;22450:10;22458:2;22450:10;;;;;;22494:11;;22598:4;22589:13;;22585:5035;;;22717:1;22707:6;:11;;;;:27;;;;;22731:3;22722:6;:12;;;22707:27;22703:501;;;22758:6;:11;;22768:1;22758:11;22754:399;;;-1:-1:-1;22778:4:0;22754:399;;;22818:6;:11;;22828:1;22818:11;22814:339;;;-1:-1:-1;22838:4:0;22814:339;;;22879:6;:13;;22889:3;22879:13;22875:278;;;-1:-1:-1;22901:4:0;22875:278;;;22941:6;:13;;22951:3;22941:13;22937:216;;;-1:-1:-1;22963:4:0;22937:216;;;23004:6;:13;;23014:3;23004:13;23000:153;;;-1:-1:-1;23026:4:0;23000:153;;;23066:6;:13;;23076:3;23066:13;23062:91;;;-1:-1:-1;23088:4:0;23062:91;;;23127:6;:13;;23137:3;23127:13;23123:30;;;-1:-1:-1;23149:4:0;23123:30;23188:1;23179:10;;22703:501;23257:11;;;23253:2161;;23317:4;23312:1;23304:9;;;23303:18;23350:4;23304:9;23343:11;;;23339:616;;;23390:4;23382;:12;;;23378:550;;23403:2;23396:9;;;;;;;23378:550;23501:12;;;23497:431;;23522:11;;;;;;;;-1:-1:-1;23515:18:0;;-1:-1:-1;;23515:18:0;23497:431;23572:4;:12;;23580:4;23572:12;23568:360;;;23593:11;;;;;;;;-1:-1:-1;23586:18:0;;-1:-1:-1;;23586:18:0;23568:360;23643:4;:12;;23651:4;23643:12;23639:289;;;23664:27;23673:5;23667:11;;:2;:11;;;;23685:5;23680:2;:10;23664:2;:27::i;23639:289::-;23730:4;:12;;23738:4;23730:12;23726:202;;;-1:-1:-1;;;;23751:17:0;;;23763:4;23758:9;;23751:17;23744:24;;23726:202;23808:4;:12;;23816:4;23808:12;23804:124;;;-1:-1:-1;;;;23829:17:0;;;23841:4;23836:9;;23829:17;23822:24;;23804:124;23886:4;:12;;23894:4;23886:12;23882:46;;;23907:21;23916:2;23910:8;;:2;:8;;;;23925:2;23920;:7;23907:2;:21::i;23882:46::-;24067:4;:12;;24075:4;24067:12;:28;;;;24083:4;:12;;24091:4;24083:12;24067:28;24063:767;;;24131:2;24126;:7;24119:14;;;;;;;24063:767;24177:4;:12;;24185:4;24177:12;:28;;;;24193:4;:12;;24201:4;24193:12;24177:28;24173:657;;;24241:2;24236;:7;24229:14;;;;;;;24173:657;24287:4;:12;;24295:4;24287:12;24283:547;;;24335:2;24330;:7;24323:14;;;;;;;24283:547;24373:4;:12;;24381:4;24373:12;24369:461;;;24422:2;24417;:7;24409:16;;;;;;;24369:461;24460:4;:12;;24468:4;24460:12;24456:374;;;24509:2;24504;:7;24496:16;;;;;;;24456:374;24548:4;:12;;24556:4;24548:12;24544:286;;;-1:-1:-1;;;;24593:7:0;;;24591:10;24584:17;;24544:286;24637:4;:12;;24645:4;24637:12;24633:197;;;24698:2;24680:21;;24686:2;24680:21;;;:29;;24708:1;24680:29;;;24704:1;24680:29;24673:36;;;;;;;;;24633:197;24745:4;:12;;24753:4;24745:12;24741:89;;;24793:2;24788:7;;:2;:7;;;:15;;24802:1;24788:15;;24741:89;23253:2161;;;;24854:6;:13;;24864:3;24854:13;24850:564;;;24900:2;24894;:8;;;;24887:15;;;;;;24850:564;24934:6;:14;;24944:4;24934:14;24930:484;;;25000:4;:9;;25008:1;25000:9;24996:51;;;-1:-1:-1;;;25025:21:0;;;25011:36;;24996:51;25076:4;:12;;25084:4;25076:12;:28;;;;25092:4;:12;;25100:4;25092:12;25076:28;25072:328;;;25159:4;:12;;25167:4;25159:12;25155:26;;;25178:3;;;25155:26;25203:8;25237:115;25249:10;25244:15;;:20;25237:115;;25321:8;25292:3;25321:8;;;;;25292:3;25237:115;;;25380:1;-1:-1:-1;25373:8:0;;-1:-1:-1;;25373:8:0;25072:328;22585:5035;;;25443:4;25434:6;:13;;;25430:2190;;;25467:6;:14;;25477:4;25467:14;25463:1046;;;25530:42;25548:2;25553:1;25548:6;25558:1;25547:12;25542:2;:17;25534:26;;:3;:26;;;;25564:4;25533:35;25570:1;25530:2;:42::i;:::-;25523:49;;;;;;25463:1046;25597:6;:14;;25607:4;25597:14;25593:916;;;25660:45;25678:2;25683:1;25678:6;25688:1;25677:12;25672:2;:17;25664:26;;:3;:26;;;;25694:6;25663:37;25702:2;25660;:45::i;25593:916::-;25730:6;:14;;25740:4;25730:14;25726:783;;;-1:-1:-1;;;25800:21:0;25819:1;25814;25809:6;;25808:12;25800:21;;25853:36;;;25920:5;25915:10;;25800:21;;;;;25914:18;25907:25;;25726:783;25957:6;:14;;25967:4;25957:14;25953:556;;;25998:3;25991:10;;;;;;25953:556;26032:6;:14;;26042:4;26032:14;26028:481;;;26111:2;26116:1;26111:6;26121:1;26110:12;26105:2;:17;26097:26;;:3;:26;;;;26127:4;26096:35;26089:42;;;;;;26028:481;26156:6;:14;;26166:4;26156:14;26152:357;;;26235:2;26240:1;26235:6;26245:1;26234:12;26229:2;:17;26221:26;;:3;:26;;;;26251:6;26220:37;26213:44;;;;;;26152:357;26282:6;:14;;26292:4;26282:14;26278:231;;;-1:-1:-1;;;26352:26:0;26376:1;26371;26366:6;;26365:12;26360:2;:17;26352:26;;26410:41;;;26482:5;26477:10;;26352:26;;;;;26476:18;26469:25;;25430:2190;26529:6;:14;;26539:4;26529:14;26525:1095;;;-1:-1:-1;;;26596:4:0;26590:34;26622:1;26617;26612:6;;26611:12;26606:2;:17;26590:34;;26672:27;;;26652:48;;;26722:10;;26591:9;;;26590:34;;26721:18;26714:25;;26525:1095;26760:6;:14;;26770:4;26760:14;26756:864;;;-1:-1:-1;;;26827:6:0;26821:36;26855:1;26850;26845:6;;26844:12;26839:2;:17;26821:36;;26905:29;;;26885:50;;;26957:10;;26822:11;;;26821:36;;26956:18;26949:25;;26756:864;26995:6;:14;;27005:4;26995:14;26991:629;;;-1:-1:-1;;;27057:20:0;27075:1;27070;27065:6;;27064:12;27057:20;;27105:36;;;27169:5;27163:11;;27057:20;;;;;27162:19;27155:26;;26991:629;27202:6;:14;;27212:4;27202:14;27198:422;;;27257:2;27250:9;;;;;;27198:422;27280:6;:14;;27290:4;27280:14;27276:344;;;-1:-1:-1;;;27342:25:0;27365:1;27360;27355:6;;27354:12;27349:2;:17;27342:25;;27395:41;;;27464:5;27458:11;;27342:25;;;;;27457:19;27450:26;;27276:344;27497:6;:14;;27507:4;27497:14;27493:127;;;27534:3;27527:10;;;;;;27493:127;27564:6;:14;;27574:4;27564:14;27560:60;;;27601:2;27594:9;;;;;;27560:60;27630:29;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;12616:509;12699:7;12718:18;;:::i;:::-;-1:-1:-1;12778:4:0;12820:2;12809:13;;;;12801:40;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;12927:13;;;;;;;:28;;;12944:11;12927:28;12923:90;;;12999:3;12971:5;:15;;;12987:8;12971:25;;;;;;;;;:31;;;;:25;;;;;;:31;12923:90;13034:12;;;;;13023:23;;;;:8;;;:23;13086:1;13071:16;;;13056:31;;;13105:13;:11;:13::i;3696:5796::-;3739:7;3758:18;;:::i;:::-;-1:-1:-1;3861:15:0;;:18;;;;3818:4;3948:18;;;;3988;;;;4028;;;;;3818:4;;3841:17;;;;3948:18;3988;4061;;;4075:4;4061:18;4057:5256;;;4127:2;4152:4;4147:9;;:14;4143:132;;4255:4;4250:9;;4242:4;:18;4236:24;4143:132;4292:7;;;4288:141;;4324:10;;;;;4352:16;;;;;;;;4324:10;-1:-1:-1;4288:141:0;;;4412:2;4407:7;;4288:141;4057:5256;;;;4449:10;:18;;4463:4;4449:18;4445:4868;;;1514:10;4502:14;;4445:4868;;;4537:10;:18;;4551:4;4537:18;4533:4780;;;4613:1;4608:6;;4533:4780;;;4635:10;:18;;4649:4;4635:18;4631:4682;;;4710:4;4695:12;;;:19;4728:26;;;:14;;;:26;4775:13;:11;:13::i;:::-;4768:20;;;;;;;;;;;4631:4682;4809:10;:18;;4823:4;4809:18;4805:4508;;;4968:14;;;4964:2075;;;;;5073:22;;;1747:1;5073:22;5069:1970;;;5226:10;5239:27;5247:2;5252:10;5247:15;5264:1;5239:7;:27::i;:::-;5325:11;5356:6;;5376:17;;;;5395:20;;;;;5356:60;;;;;;;;;;;;;;;;;;;;5226:40;;-1:-1:-1;5325:11:0;;;;5356:6;;;;;:19;;:60;;;;;;;;;;;:6;:60;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;5356:60:0;;;;;;;;;-1:-1:-1;5356:60:0;-1:-1:-1;5567:1:0;5559:10;;5657:1;5653:17;;;5728;;;5725:2;;;5758:5;5748:15;;5725:2;;5837:6;5833:2;5830:14;5827:2;;;5857;5847:12;;5827:2;5959:3;5954:1;5946:6;5942:14;5937:3;5933:24;5929:34;5922:41;;6034:3;6030:1;6019:9;6015:17;6011:27;6004:34;;6154:1;6150;6146;6134:9;6131:1;6127:17;6123:25;6119:33;6115:41;6277:1;6273;6269;6260:6;6248:9;6245:1;6241:17;6237:30;6233:38;6229:46;6225:54;6207:72;;6400:10;6396:15;6390:4;6386:26;6378:34;;6512:3;6504:4;6500:9;6495:3;6491:19;6488:28;6481:35;;;;6608:33;6617:2;6622:10;6617:15;6634:1;6637:3;6608:8;:33::i;:::-;6659:20;;;:38;;;;;;;;;-1:-1:-1;5069:1970:0;;-1:-1:-1;;5069:1970:0;;6759:18;;;1666:1;6759:18;6755:284;;;6940:2;6935:7;;6755:284;;;6986:10;6981:15;;1822:3;7014:10;;6755:284;4805:4508;;;7059:10;:18;;7073:4;7059:18;7055:2258;;;7222:15;;;1593:1;7222:15;;:34;;-1:-1:-1;7241:15:0;;;1628:1;7241:15;7222:34;:57;;;-1:-1:-1;7260:19:0;;;1705:1;7260:19;7222:57;7218:1376;;;7304:2;7299:7;;7218:1376;;;7374:23;;;1790:1;7374:23;7370:1224;;;7453:10;7466:27;7474:2;7479:10;7474:15;7491:1;7466:7;:27::i;:::-;7565:17;;;;7453:40;;-1:-1:-1;7733:1:0;7725:10;;7823:1;7819:17;7894:13;;;7891:2;;;7916:5;7910:11;;7891:2;8190:14;;;8004:1;8186:22;;;8182:32;;;;8083:26;8107:1;7996:10;;;8087:18;;;8083:26;8178:43;7992:20;;8282:12;8348:17;;;:23;8412:1;8389:20;;;:24;8000:2;-1:-1:-1;8000:2:0;7370:1224;;7055:2258;8614:10;:18;;8628:4;8614:18;8610:703;;;8712:2;:7;;8718:1;8712:7;8708:595;;;8797:14;;;;;:40;;-1:-1:-1;8815:22:0;;;1747:1;8815:22;8797:40;:62;;;-1:-1:-1;8841:18:0;;;1666:1;8841:18;8797:62;8793:376;;;8888:1;8883:6;;8793:376;;;8930:15;;;1593:1;8930:15;;:34;;-1:-1:-1;8949:15:0;;;1628:1;8949:15;8930:34;:61;;;-1:-1:-1;8968:23:0;;;1790:1;8968:23;8930:61;:84;;;-1:-1:-1;8995:19:0;;;1705:1;8995:19;8930:84;8926:243;;;9043:1;9038:6;;8926:243;;8708:595;9212:10;9207:15;;1856:4;9240:11;;8708:595;9323:15;;;;;:23;;;;:18;;;;:23;;;;9356:15;;:23;;;:18;;;;:23;-1:-1:-1;9401:12:0;;;;9390:23;;;:8;;;:23;9453:1;9438:16;9423:31;;;;;9472:13;:11;:13::i;:::-;9465:20;;;;;;;;;3696:5796;:::o;10739:1393::-;10829:7;10848:18;;:::i;:::-;-1:-1:-1;10908:4:0;10931:10;10963:4;10955:12;;;;10951:984;;;-1:-1:-1;10989:8:0;;;;10951:984;;;11034:4;:12;;11042:4;11034:12;11030:905;;;11062:13;;;:8;;;:13;11030:905;;;11112:4;:12;;11120:4;11112:12;11108:827;;;-1:-1:-1;11146:8:0;;;;11108:827;;;11191:4;:12;;11199:4;11191:12;11187:748;;;11219:13;;;:8;;;:13;11187:748;;;11269:4;:12;;11277:4;11269:12;11265:670;;;11405:9;11356:16;11337;;;11356;;;;11337:35;11412:2;11405:9;;;;;11387:8;;;:28;11429:22;:8;;;:22;11265:670;;;11472:4;:12;;11480:4;11472:12;11468:467;;;11554:10;11541;;;11554;;;11541:23;11604:2;11597:9;;;;;11579:8;;;:28;11621:22;:8;;;:22;11468:467;;;11664:4;:12;;11672:4;11664:12;11660:275;;;11747:2;11729:21;;11735:2;11729:21;;;;;;;;11711:40;;:8;;;:40;11783:21;;;;;;;;;;;;;;11765:40;;:8;;;:40;11660:275;;;11826:4;:12;;11834:4;11826:12;11822:113;;;11890:2;11885:7;;:2;:7;;;;;;;;11874:18;;;;:8;;;:18;11917:7;;;;;;;;;;;;11906:18;;:8;;;:18;11822:113;11949:13;;;;11945:75;;12006:3;11978:5;:15;;;11994:8;11978:25;;;;;;;;;:31;;;;:25;;;;;;:31;11945:75;12041:12;;;;;12030:23;;;;:8;;;:23;12093:1;12078:16;;;12063:31;;;12112:13;:11;:13::i;:::-;12105:20;10739:1393;-1:-1:-1;;;;;;;10739:1393:0:o;15271:1320::-;15358:14;15375:23;15387:10;15375:11;:23::i;:::-;15358:40;;15444:1;15438:4;15434:12;15431:2;;;15459:1;15456;15449:12;15431:2;15579;15760:15;;;15597:2;15587:13;;15575:10;15571:30;15568:1;15564:38;15719:17;;;15518:20;;15704:10;15693:22;;;15689:27;15679:38;15676:61;16019:4;16016:1;16012:12;16197:1;16182:319;16207:2;16204:1;16201:9;16182:319;;;16322:2;16310:15;;;16263:20;16353:12;;;16367:1;16349:20;16386:42;;;;16450:1;16445:42;;;;16342:145;;16386:42;16403:23;16418:7;16412:4;16403:23;:::i;:::-;16395:31;;16386:42;;16445;16462:23;16480:4;16471:7;16462:23;:::i;:::-;16454:31;;16342:145;-1:-1:-1;;16225:1:0;16218:9;16182:319;;;-1:-1:-1;;16521:4:0;16514:18;-1:-1:-1;;;;15417:1168:0:o;13131:514::-;13418:19;;;13441:7;13418:31;13411:3;:39;;;13511:14;;13558:16;;13552:23;;;13544:71;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;13625:13;13131:514;;;:::o;-1:-1:-1:-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;;;;:::o;:::-;;;;;;;;;;;;;;;;;;;;;;;;:::o"

func init() {
	if err := json.Unmarshal([]byte(MIPSStorageLayoutJSON), MIPSStorageLayout); err != nil {
		panic(err)
	}

	layouts["MIPS"] = MIPSStorageLayout
	deployedBytecodes["MIPS"] = MIPSDeployedBin
}
