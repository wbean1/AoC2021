package day8

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
func SortStrings(s []string) []string {
	sorted := []string{}
	for _, str := range s {
		sorted = append(sorted, SortString(str))
	}
	return sorted
}

func Run() {
	input := Input()
	count := 0
	for _, value := range input {
		split := strings.Split(value, " ")
		for _, word := range split {
			if len(word) == 2 ||
				len(word) == 3 ||
				len(word) == 4 ||
				len(word) == 7 {
				count++
			}
		}
	}
	fmt.Printf("part1: digits appear %d times\n", count)
	sum := 0
	for key, value := range input {
		sum += DecodeInput(key, value)
	}
	fmt.Printf("part2: sum is %d\n", sum)

}

func DecodeInput(key, value string) int {
	inputParts := strings.Split(key, " ")
	inputParts = SortStrings(inputParts)
	inputMap := make(map[string]int)
	reverseMap := make(map[int]string)
	for _, part := range inputParts {
		if len(part) == 2 {
			inputMap[part] = 1
			reverseMap[1] = part
		}
		if len(part) == 3 {
			inputMap[part] = 7
			reverseMap[7] = part
		}
		if len(part) == 4 {
			inputMap[part] = 4
			reverseMap[4] = part
		}
		if len(part) == 7 {
			inputMap[part] = 8
			reverseMap[8] = part
		}
	}
	for _, part := range inputParts {
		if _, ok := inputMap[part]; !ok { // only care about things we don't already know
			if containsAllChars(part, reverseMap[1]) {
				// is either 0, 3, 9
				if len(part) == 5 {
					inputMap[part] = 3
					reverseMap[3] = part
				} else {
					// either 0 or 9...
					// if all the letters are in 4, then 9
					// else 0
					if containsAllChars(part, reverseMap[4]) {
						inputMap[part] = 9
						reverseMap[9] = part
					} else {
						inputMap[part] = 0
						reverseMap[0] = part
					}
				}
			}
		}
	}
	for _, part := range inputParts {
		if _, ok := inputMap[part]; !ok { // only care about things we don't already know, 2, 5, 6
			if len(part) == 6 {
				inputMap[part] = 6
				reverseMap[6] = part
			}
		}
	}
	for _, part := range inputParts {
		if _, ok := inputMap[part]; !ok { // only care about things we don't already know, 2, 5
			// if all the letters are in 6, then 5
			// else its 2
			if containsAllChars(reverseMap[6], part) {
				inputMap[part] = 5
				reverseMap[5] = part
			} else {
				inputMap[part] = 2
				reverseMap[2] = part
			}
		}
	}
	fmt.Print(inputMap)
	valueParts := strings.Split(value, " ")
	valueParts = SortStrings(valueParts)
	valueStr := ""
	fmt.Println(valueParts)
	for _, part := range valueParts {
		valueStr += strconv.Itoa(inputMap[part])
	}
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(valueInt)
	return valueInt
}

func containsAllChars(big string, small string) bool {
	for _, char := range small {
		if !strings.Contains(big, string(char)) {
			return false
		}
	}
	return true
}

func Input() map[string]string {
	input := map[string]string{
		"cgdf eagcbf fc adefg eacdb fbedga geafcd efc dacfe fdgaecb": "dcefbag dgcf fc daefc",
		"bdecf dcagb gbf gcbdf deacbf fg fdebgc fegdcba dgef bgefac": "dbfec gbacefd gf bfg",
		"cfeag becgda bag ab abcd ecgdb gdefba agcbe gdcebf fgebadc": "ab cdegbf cbda cgfebad",
		"gc bcfgea ebdcf cbedg edbfac ebfdgc bacfged ceg dfcg abdeg": "eacbdf geabd dcbef fcgebd",
		"fgbceda afdbge fbcad badgec cfde gfcba ebcda afd abfdce df": "fad fdcab fda befcadg",
		"gfacd dcf bacedg afgec afdbec df agcdb agcdebf gfbd cdgbfa": "df df bceagd ecafg",
		"gbdfe ag gda bfadcg gbedfa gdefcb beag gedaf cdaef bdgcfea": "gacfdb gedfab ga gda",
		"deabg cag cfdea bfgcda ecgb fcbedga cg dcage agebdc gdbfea": "bacefdg bdcgfa adcgfb afedc",
		"agfecbd gbcadf ce defc dabcfe fgeab gbdcae ecb fcabd bcafe": "ebafcd ce ceb ce",
		"fa adfebc gdacef bfeca begca bfad bdfce febcagd dfgceb fea": "afe cfdgbe af eabfc",
		"abge eb efbgadc cadeg cabefd gbced cgdfb abegdc acfged bce": "dgeac ebdagc edagc gacbfde",
		"cfbedg gdafbe fgebd bc cbgdae bcagfed ecgfa bcfge bgc bdcf": "gafebdc bgfdce bcfged dfgbe",
		"cafebgd gafdc bgcef bcagf fba cabefd ba baeg feabgc bdfcge": "fba ab ab gaeb",
		"ecgfbda afbeg fcdeg fecbga cea afbc abdegc befagd ac gcaef": "beadgc eca cae gebacfd",
		"cedag gf afbcd cbaedgf eacbdg acfegb fcadg fgedac gfed fgc": "dacbf daecgb bagcef cgf",
		"fagbd gdaceb ebf efbad gacebf fcde ef ceabdf aecgbfd dbeac": "dbeac bagfd cbfdae fbe",
		"begadcf aedb cgefa ebdgac acfdbg cdebfg agebc ba bac gdbce": "ab faceg bac adbe",
		"bgde gb abefcg fdaecg gcbdf ebdgafc dfgce gfebdc cgb acfdb": "fdbcg gb dcegaf cfdge",
		"efbagd gefcdb dgbefca badce cegdb efdgb gfcbae cge dcgf cg": "dbeca bdaec ecg egc",
		"begdca caef gfebd cgbefda cgf cf gdecf cfagde cagbfd agdce": "dgfcbea fcae abcfgd acdgbf",
		"cadfbeg acd egbad dcbf efcab cd agfbce daebc gefdac afbced": "bdefacg cd abgfec dgecfa",
		"fd bfd bgcfd bagcfe edgf fcbge acbdg cdebgf efbadc bgeacdf": "df bdcga fbd gbedfc",
		"fdgbca deca de edagbc ebadg gbfea bcgda bcdegf dge gcebadf": "abgcfde de gdeba acde",
		"df gdbecaf bdfe fbcgae fedabg fdaegc cgdba fdg adbgf egbfa": "gdf fbgaec egafb gadbf",
		"egfbd cefg gfdbac fbdea bedcg gadcfeb gdefbc gf bfg adgcbe": "dfebg dbacfg gcafdb dbafe",
		"bgadc bacfdg acbdge egafdcb ea deca cfbaeg eab dabeg fgbde": "aecd bdcag dgeacb bafdecg",
		"cfbeag febcgad cge dacbg facdbg ge gcbde eadg acbgde fcbed": "cagbed gcabfe gdcfbae bdgec",
		"cgefa eabf gbedc gcefb cdbfaeg dfabgc fb beacgf gfb gfedac": "bf afceg bfg fgaec",
		"fdg gd edfcbg gaed gfcda febcga ecdafg cgfebad fgcae dfcab": "cafegd cedbagf gfcea aecgfb",
		"dcafbe ebcadfg adcgbf abfcd bg fdaeg agfdb gcba bgd efbgdc": "cgab ecbdgf agdbf efcbgd",
		"adcbf dagfcb dfecb dagc feadbg ad bgacf cegdbfa bad gcafbe": "cfdbage deafcbg fcagdb cedgbfa",
		"afbecg fe egdcf gef afde acdegf feadgbc fdbgca gdceb fgdac": "edcbg edafgcb afgcdb dbgce",
		"gbdfc aecdgb ef bfe acdbe aefc fceadb bcaefgd agbefd bfced": "fbe gcbdf fcea adbec",
		"fcgae ebcfa becdagf fgadcb cba bfdea bedc ebacdf bfgead bc": "egbfadc cb cb caefg",
		"efbd abcde cfeag aecbdg daf df eafcd gcafbd fdeabc cgbefad": "gbefadc fecga cgfae acdbe",
		"afegbc becaf gaf fcagbd egbf adbcef gafec aedcg afdcegb gf": "bcfae efacb bcaef fga",
		"gf bgecdf gaebfdc fdcbg fdg ebfg agcdb cfdbe febacd efcgda": "fdcbe gbcdf bcdga cgafde",
		"fdebac dgac dfega dec feadbcg fdcgea cefgb dfaegb egcdf dc": "decgf gdca ecd cfgaebd",
		"fabdgc dceab dbfgc dagfeb begdfc bdgca dabfcge dga ga facg": "gfacdb adgcbef bgedaf facg",
		"cfeag cgdfa aegb fecdgb dfacbe dcgbfae ebfca ge egc gcbafe": "ebgfdca aebg gcebfd cbfea",
		"dcafg edcgba afebd faedbgc bgef edafbc fabdg bgefad dbg bg": "ebdagc abgefd afcebd bfeadc",
		"dfbgce efcbg efgbac fcbea abdfgce fecad bgac fba ba dfaebg": "cafdgeb bcag adcef ab",
		"acfb bec fabged gbeafc ecgabd fecdg gabfe bc cbgef gcbeadf": "cbaf bafcdge baegf fagbde",
		"fgebcad cgadfe agbdcf gfadc edfcg badefc efag ef gdbec efd": "dcfeg becdfa bdecg bgdec",
		"bafdgc fdcegba deafb dcbfe abf gedfa ab ecab cbaedf fecdbg": "aecb ba baf dbefc",
		"bgfeca cadeg adfgc gea ae bcdagef edba abdceg gbdec fcebdg": "dcgefab efcbag fcdag beda",
		"fcag afecbg cabdge dbefa cef ecbgfda begdcf fceba fc gbaec": "gfca bdcegf cfe aebcfdg",
		"cfbga df bgaed dgef eafdgbc cdabfe bdecag adgefb dfa agbdf": "fd gdfe df df",
		"edagfb cdfbaeg aecd gcbad dcgbef dbeagc da dcegb dab fgcba": "cdae ecad abgdfec cfgab",
		"ebc agbcfe eb bdfcag decgbaf agecd egdbfc egacb afeb bacgf": "ebc cbe abfe fegbcd",
		"dcefabg bfaed feg ebgfa eg egad cgdebf cgfab degbfa acdefb": "edbfgca cdfaeb gfe feg",
		"adfebc dbgce ecfdgb ba daegb beadgc dfgea dcegfab agcb aeb": "cagb acebfdg cbga eafgbcd",
		"afcbged gfdbc cd bdc edgfbc befacg edcf abdceg dgbfa cfegb": "gfcdb gebcf bcd fdce",
		"gbfceda bf bfd cedgf fgdeb gfab befgad gcebda debacf gbade": "gfdbe fdebag adegfb gdeba",
		"eacd cdegafb bac gbedcf gcebd afgdb ca gadbc gdcbae eabfgc": "becgd fcbgea aedc edac",
		"bagd bfgac adbfc aefcg gcbfad cadfbe dgecfb gfb gb gefdabc": "fbg dcegabf bgad bdcfa",
		"bfaedg bfcga cdgfe fcdgba gabcfe eacb debfcag eb fbe gcefb": "bef efdcbag cbfaegd ebca",
		"afcebg dbfag dabe dgbef cebfgd gdafc egadbcf ba abg afdgbe": "bga bga gab ba",
		"ecdabg defgc bdacgf dae eadcg cdbag gcbdafe aecb adebgf ea": "bfedagc adcbg eagdc bdgeaf",
		"acgde bcdea bfcd bface becfad afdcgeb dfaebg bd adb cfgeba": "eacdb cbgefa fabcge bcdea",
		"gecfad fbcegd dbfaceg fgabe dge gfbed fbdgca gdfcb bdec ed": "gdbfec egd eagbf gadcfe",
		"dagfe cbafdge bag dgebcf cfab dafcgb adgbf ab fgdcb geacdb": "gacfdb fgbcd gafdb afbc",
		"afecbdg bdc adefb eabdcg fcged bfdce bfgead bdafce bc bfca": "dbcega bc dabgefc fcba",
		"cdgafb bafcge fb bdfg dagcb cbgafed fcdba cedfa cbdgea fba": "deacf dbecga fdbacg gabfdc",
		"dcf bcadf ebgdfa bcadgf bedac fc dfcbeg beadfcg afgbd fagc": "dcf bgdfa cbgdfe agebdcf",
		"gadbf egdba gfdeca aeg cegdb gdbcfa faeb ae gdaefcb gebfda": "ebdag eafb baged ae",
		"ebfgc daeb cae gceba cgabfde cgabdf ea abcgd bdaecg ecgadf": "ace eagbcfd ea efacgd",
		"bdeaf ag bgdeca dag agfdb agcf dcgbf cfgbeda dfcbeg cdagbf": "fbaed fgdabce beacdg ecabdfg",
		"ecagf gcfdea ba gbfa badgce cab abfce fecdb bagcfe bafegcd": "faecdg ab bedcf edcfb",
		"geadf adgebf eda cdgfab bedg bfadec de gefbdca adfbg gafce": "ecagfbd cafge de eda",
		"fb agcbe gbfc cbaegd ecgfdab edafg bdcefa efgba bfa gabecf": "bfa fcbade gceab fba",
		"bdaefcg dgacfe fbgadc gfbea dcae aegdf ad gcdfe bgdcef dga": "cdae da cade gdfcae",
		"cbeg eacfg fbcage acfdge bgeafd gbfacde afcgb bga gb cbdfa": "bg gabfde gba fdagbe",
		"bca defgab bc gfcae gaebd dfebac bcage deafgcb gcebda cgbd": "edbgfa cb gfcea cb",
		"fcb cgadfeb bcagf abdgf cegdfa begcaf gcbe fcebda cafeg cb": "cfeadbg cfb dfbga gdaecf",
		"fdebag abdcge gfdb efbac gcefad eagfb agb degaf bg acfegbd": "agefbdc gfdb gb bedafg",
		"ec dce cgbaed egabdf dbeacf cagdf ebcg agcde gbedacf eagbd": "bafecd eadgbc cdgfa degac",
		"beg fageb bfgad ge abecf bdcfae caeg efadgbc egbfdc cgebaf": "dfabg fabgd fbdceg beg",
		"gdcfe bacfeg dfb db cfbdae bade aegfcdb fbecd bagdfc fceab": "cbagfe bd bd bd",
		"gab cagdfb fgecdb fgaed ab gcbdf baecgd abcdfge gbfad bacf": "gbdfca gab bgdcea abg",
		"ecfga afdbeg gdceafb gcedf ecgabf afc dcgfab ca bcae gbaef": "cgeadbf fagec cfa aebfdcg",
		"fgaebc aefcdg fdceb dfbeg gbf agefbd fegdbac bdga bg fgdae": "bcdfe fgb acfgbe gdab",
		"dgaecf fadbgc cdfeab cagbe cbd bd bdeac gadbfce afced bdfe": "gfdace febd db cgfade",
		"bfdgac efgdab abd edcbf baeg ab afegd dagecf cgbfdae edbfa": "gdaef dbefc dab ba",
		"acgb bda gdaefb eabcfdg ba abced gdcbe dacef gcdeab bcdefg": "adefc bcade deagcb adb",
		"fgabc cdfabg cga dgcfeb deagbc ag fgda dcfagbe caefb fgbdc": "dagfbec bfacg gfdbac edgcab",
		"gcfdeb fbgdca beagf fgbad ecfba efg egbfacd ge daeg defgba": "fcbgade fbgae fbcae adgbfe",
		"abe dgab defca dbeacg ba fgbdce cbdeg bfceag cebagdf eadcb": "bgda dfgbce dfegbc gbcfaed",
		"cbafg efdgcb cfaebg afecb aefg gafcdeb ecf ef becda fdgcab": "cafgeb acgbf acebgf dcbfag",
		"eb cebg dbe dceab dcbage ecgadf gcaed cbfda agbdef gadefcb": "agebdc baedc dacfb gcebda",
		"degcb cagbef dfceba gbf fg gbacdf gfeacdb fbgec fgae ebafc": "bfcdga bfg fbaec fbg",
		"dbgaf bca gacde cefbga acebgd edcb gbacd fbcagde cb cdefga": "gcdea fgcebad cab cdebfga",
		"bdgfea gb fcdeg gfcdae ecgfbd cdgb egcfb ebafdcg befca gbf": "gb dagebcf dgfec dfgcbe",
		"geacb fg badfc dagecfb bfcgad gfb fcbga gcdf efadbg befadc": "dfbgcae badgcf dcbfa aebcg",
		"febdc aecfd aebgfd gcdfbe dbgcf bef adbfcg dbefacg gbce be": "feb eb ebdcf dceaf",
		"acde cdegb decfbga gebac eag cegfbd ea cdbgea ebgafd cbagf": "ae gea bgdafe cagbe",
		"gefac egbfadc degbfa badgce gef fe efgbac agecb fcdag bfec": "afdegb adbecg agbcde agfce",
		"gfceb bg edcbf afcegd gacb fagec fegdbca gbaefc fgb edgafb": "bgf agcb eagcf facged",
		"bdcagef dcfbg ceafd bead adfcb fgeadc cba cefbga ba cfdeba": "gcabfed decfag ecgdfa fgbdc",
		"cdbfge bfgac efgad cbae daegbcf fgaecb ebg begfa be gdacbf": "eb be efbga adefg",
		"gfabd bgaed gf cebfga fgb eafcbdg degf aegfbd acdfb gdbeac": "abged aebdfg bfgeca bdagec",
		"gcafd efagcd edcg bafdeg bfgdeac ebcgaf gd facdb dgf efagc": "fbadc fbcda cbeafg cdebgfa",
		"dfeba cfb dbcgef dbceafg efgca cbeaf fecdag fbgcea bc gacb": "aebcf fcb adgfce fcbea",
		"badgf agdcef dafbec ag gabe dfeab dgfeab bdgaefc gfa gdbfc": "eadfb dcefag cafdebg edfagc",
		"bef eb abec fabgec facbdg afbcg defbag dcfge fcbge deafgbc": "cgebaf eacb dcagbf bfcga",
		"bdfe dgebcf bcgefa ed cbeadg gefdcba fbceg dcgfe facdg dge": "dcfge de cgdabe dge",
		"gbadc dgbefac cgbafd dfcg dfacb ecgabf cg gcb bcfdae bedag": "cbdga cdbag gdabfc bdeag",
		"bagecd cf fbedcg cefa cafbd bfc cbedfa gafbd baecfgd caedb": "gfdebc cdeab cfae cfb",
		"fgdab efbacd decb cgaefb dbcaf acfdge fbc fdaec edgabcf cb": "edbc bc fgecab cfgade",
		"dcbga adfbc df decfab fbed cfeba afd egfcadb dcaegf befcga": "ebfd bgafec aebcgf ebdf",
		"cbdaegf fdegcb cgbd dafecg dbfgae dfbce gdfec bdf bd eafcb": "cbdg dcgb gdfaec ecabfdg",
		"bgfade bdeagc befcd daecgfb efb abecd eacbfd gdcbf cfae ef": "cfea ecdba agdcbfe gbaefcd",
		"fagd dcafbe gdabfec df bgcfe cgfdba bgadce agdbc fbd gbcdf": "bgcdf df edbcfag df",
		"gfceb ae eac adbegfc gadcef dcbfge afceb adcbf gfbeca ageb": "cbegfa cfdab ae ae",
		"adgeb fcbdg fbdgec cbged ecb ce efgabc dgefcab bfcgad dcfe": "abdegfc fgbcae cdbfag gfdbc",
		"dfbegc cb dgeac afedbg cebf ebacdfg dfegb gdbce cbg fgbdca": "efagdb gdbecf dbgec bacgfd",
		"fdcgabe adfceg fegbcd bg cdefg bfg cegfb ebfca bfagcd edgb": "gbf dbge fcedg gb",
		"dgefc cdaef cbafgd dbgecaf egfb bcegfd gabced eg bdfgc gde": "acgdbf bedgac dfgce adcef",
		"fgeac dgbfeca gaecdf cb bcegfa cgbf ebc cdageb defba ecbaf": "fbace ecb gface gbecad",
		"abedfg cdefb gfb egafd bg ecagfd gbfde agcfbe bdga dgceafb": "fbg efgbd bfg bdfce",
		"gbfcde fb ebdga faced bfca eadfgc adbcef eabdf gedfabc bef": "bfca bf bdefa fb",
		"fbgadec gcbe bfgca cg bacfde egfdca cag afbec gadbf cagbfe": "gbecafd bacdfe cga egcb",
		"cfead fcdbge ebf bace eafdcg bagdf fgbecad be eadfb fecbda": "efcdab feb aefdb ebf",
		"dabfc cfeabd ec cae cagbdf fadeg cdafbeg dbecag bcfe efcad": "bcef aec eac bfacd",
		"fed fbdgea adgfce gfcd acdge acdef afbec df gcabed bdcefga": "aebgfd afecd bcfae ecdga",
		"fcabe da deac fdgeb fbead facebg gcdbefa fcabgd dfbeac dab": "fecab bcfgae fcaeb ecbgaf",
		"cdebg dgebac abe fagcbed ea cbagf adeg cfdbge gacbe edbcaf": "dcebaf fgbca adeg cafgbed",
		"gedfcb bg bgc gfacd fdgcabe dgba bfagc afbec bgacdf eadcfg": "gfcab bcg bg dgba",
		"fecga fgbeda fadec fdcbega cdefba becdf dgfbec ad dacb dfa": "cedbfg gebadf agefdb bdac",
		"bcgade agefd ebfg bag gfdbea afegdcb cfdba dagfb bg gfcdae": "adfbg bga gb agb",
		"fgaced bcdag dcbagef gedcbf dcfge fga af efda gacdf eabgfc": "bdcga fga dcfeag fead",
		"dega abd adbgc eabdgfc dbcefg da fadbec ebcgd cafgb adgcbe": "cbdga cabdg adgefbc bfcgde",
		"bfdacg eabgc bce abgdc cebgafd be ecadgb faebdc bdge gface": "bdeg acefg gefbadc dcgba",
		"fedgc cadefg cedagbf adcfe abfdc efa ae agfedb gace bfecgd": "acbfd ae gcbafed cdabefg",
		"egdcab egdfba gca abfec bafgd gcbefda gfbcad cg fcdg bcfga": "gc gc edabgc cga",
		"eb gdbfca befdc gdcfb ebdg efb edcbgaf ecfdbg fedac egbfca": "gbdcf acfegb aedfc dafcebg",
		"aebcgd bacg ecgbafd dafbec ga dcbae aedbg gad bdgef fadgce": "adegb gad becagd edbcfa",
		"cgad ceabdfg fedbc gd cfbage gfcba abgefd fdabgc dcgfb gdf": "dcgfb gbdafec gbfadc dgca",
		"acgebd fdebga cde cedfa gfceabd baefd cfeb cgadf ce fedcab": "afedb deabf afdce fabged",
		"gcabe adge bcdfea cgebf bgcfda aeb gdacb ea aegdcb cefagbd": "bea gaed agbce dgfacb",
		"cefad cadgbfe cgeafd eg agdbfe acgbf fge cged gfcea ebfacd": "dacebf dceafb fabecd gcde",
		"gbfcde fcg ecgadb cbdge eagdf edgfc bfdc bfecag fc dbcagfe": "cfbd gfdecb gfc efdga",
		"gebafc bcfae fcebd gfabed bea bfagc fdcgba ea ceag gacbefd": "ebacfg egac bacfg febac",
		"ed gfde dcgea cdagf ead cfdgaeb gcdafb aedgfc dbaefc ceagb": "dae eda gfabcd gdbfca",
		"cfbgd bfaecg dgbecf fgcdab dbgaf fba bdcefga af gebad cdfa": "feadbcg fcad bfaegcd af",
		"gbd eabfgc egdfbc dcge dg bdcgfae dbacf cdfbg fbedga cefbg": "ecgd gbd dgce dbg",
		"gac gcdaebf agfcbe gdebcf ga bcfdag adcbg bfgcd dafg cedab": "gcafeb agcdb cbfdag ag",
		"dfg bcfaegd dg gacef dgbefc dbcaef gfcbad ebdg efcdb degfc": "efcga dfcgbe gbde cefdb",
		"bcae aegfb egdaf gdcbfa eb fbcgae afgcbde ebg dbfcge bcfga": "cdbegf ebg fbacg egb",
		"ceafgb ab edbgfac aedcgf abfc fgaedb eab bcedg ebagc acegf": "fbac ba ecagf abcgfe",
		"geafd bgeacfd bcfgda bdca cebfga gdbaf fgdcbe dgfbc ba bga": "efagd gfbda cebafgd agfde",
		"gacfed agc cfgdbe ca ecfdg fagdc agfbced bacdeg bdfga eafc": "egdfc badgf egcfbd bfadg",
		"eca ac cadgef eacbdg fecgdba dbace gbca dfgbec afedb bgcde": "gbecad bcag dcaeb cae",
		"gfaebc fg fcaebd facbe fgeab cbdafg aebgd gfce caefbgd afg": "gf cfbea edbag efgc",
		"fgaed gacedf fbdcga ecfa degabf cga ac efabgdc gecad begdc": "efca acg edbgc cag",
		"afdbgec gcebad fecgad ecagd fdac af bfaegc fae efadg efgdb": "acdge afegcb gafde fa",
		"eadcf dbecgfa dfaegb gdcb ecg adbegc gc geacd ecabgf edabg": "gdbace bcdg ebagfd cbdg",
		"fgade dgb gcbe gafcdb fecdb bfacde ecgdabf bdegf febgcd bg": "cdgfab cgbe gdb gacdebf",
		"dcfbe dbfeag dgebca fbagcde gbafdc cagf bgf dcgfb fg gcbda": "abcfdg bfg gcbfd dcfgab",
		"bgf dgcbef acbefd egadcbf bafdge bfdae bg eagb acfgd dgbfa": "deafb bdfae dfagc cdgaf",
		"bafcd egcbdaf dg cfbgde eagd gfbea bgd adbgf gdeabf befgca": "badcf fbaeg dfgab fdgeabc",
		"gea fecdag gadbfe aegdc acfde dbcefa ge bcagd bgdacef egfc": "ecgf defac gea defca",
		"bfgdae gefabc fd edgcb gefab abgcdf aecgfdb fdea fgbed dbf": "fbd cgedb fbd df",
		"adbgec bc dabfge dafeb dcb cegfd cafb becfd cgdebaf bdecaf": "febcd cbfa cb adbcgef",
		"ecgdf fbdgcae acfbg fdgaeb adf afgcbd dagfc ad adcb acgbef": "dbaefg cbfga cabgf gfacd",
		"cefa ae adcbg gefbad eagfcb decbgf aeg gebcf gfdacbe abgec": "cefdgb fedbcg cfeabg gbcea",
		"ebfag fdbcge gdbaf fcdag fgbace bfd abde gfcdeba db gefbda": "bdgecf ebad deafbg fbaged",
		"bgecd edcba bgacef dgacbe fdcabeg fbeacd gbc fgdce dgab bg": "decab bcdeg fcgbea caedb",
		"bc abged abgfdc dgbac edfgca cfba gcdbef dafgc dbc cbgdfae": "adebg bcefgd cabfedg cdgafb",
		"gcfad fb bceadf cgaeb agfbdc fbc gfaced acdbfge cgbfa gbfd": "afedbc cadfg gcbae ecbag",
		"cgadb begfcad gdcbfe bdegc edab cabfg cagfde dca ad egbdca": "adc da adc da",
		"gfcbd cbegfda dbga da bgdcef bgdcaf gceaf cfebda fdgca acd": "dcgfa fedagbc bgcfd afgdc",
		"afge ge beg abcgd afgcedb fgcabe fcbae cdebfa dgebfc cabge": "acefbgd geb fage eadcfb",
		"fgecadb fcae edgfcb aedgf gdfce eagdb af dagcef fcdabg adf": "adfecg ecaf dbafceg cgefbd",
		"gabfd bedag abegcf df fad abcgfd bafgc bfgcdae fcdb gacfed": "aefgbc agdcfb cagdfe afd",
		"degbafc ga efcga age fdcgae agcd dgfeba dgefcb fcgde cbaef": "eadfgb afcbedg degfcab fabecdg",
		"beg aefgbc eb efbd cbagd gaedb aegcfd afdge dafecbg eagdbf": "dagcb geb febd cefdga",
		"da cagfd fbgcae fgbcda gfbcaed edfcg fad ebfcad dabg bgfac": "fcgde daf cgfabe afd",
		"gefdc ebafgd fgacbd agbc fdacg dcabgef ac cda bafdec abfgd": "gacb gdfca fedcbga fdcga",
		"gedfba bgdca dfgbc bcfgde bdace bacgdf cafg ga dag bdcafeg": "debcfg cgbad acfg ecdgbf",
		"fagde fdgeca agc gc afedbg ecgf agcedb dfaecbg acdfb fgdac": "cg adfeg fdgae efcg",
		"dgcbe fe abcfg cfgead gdcbfa efc egfbadc bgacef fbae gfbce": "fbcag abef faeb cafbdg",
		"beadfc eb gbef fbgda cfbgda gadbe gdace gefadb edb cgdabef": "aedcg deb bde dgbaf",
		"afbdge dbgfcea fadcbe cb febcd bcd gfdce cfagbd bcae bdfea": "bdc gfadeb cebdf afdbe",
		"degca adbfgec bfcgda fa eafb fac bafecd fedbc cfade bgecfd": "cbdafg edbacf agcfedb fa",
		"fabde egdacb bfd fdgbac ebcda bf bdecfa fadge dfegabc fcbe": "feabd fdb cfdeba bf",
		"aedbfcg fadbg cefgba ecdfga efg ge adefc dfebac afdeg dceg": "bdfeca gcde aedcfb faecd",
		"dgaecbf cefdgb bcgdfa edb ed beafc gfed bgfcd dacegb dcfbe": "deb gdbcfa cfgdeb gfecabd",
		"bdagfc bf gcdaeb dbgac gdfce cbdfg abgf bfc abfgdec bdacfe": "dcgfb defgc bfga bf",
		"gef bgcf dafbec bcfde decgf bcefdg fg bgcfade cedag fdbgea": "bacefgd gcbf fg dbfagce",
		"fagcdb cf bcafed cefa bdegf cdaeb edfcb ecdgafb daegcb bcf": "dcfgba eagcbdf caebd afbdcg",
		"ecd dcebaf bgeafdc adfbcg gdafe cbagde ecbf fceda ce fdacb": "bgecad decfabg ec bfce",
		"begda gfbd fbgade agdbcfe gfe gacebd bfega gf bcfae adcgef": "efg fcabe geadcf abfcedg",
		"cbfga ef efcgb befcga fcea fge cbdeg dgebaf fcbgad afegcbd": "fadebg ef abefdg fe",
		"dagbc dab fbdg cdfag efcbda gecba bdcefga gdbcaf cdefga bd": "feabdc dgabc db cadgef",
		"ge beg afcbed edbag fbgcea dgaebcf adbce fbadg bgacde gdce": "gced egdc feadcb eagcdb",
		"geafbdc ca gdafce adbeg cfab bagdc dcgebf adc gdfbc gbcafd": "cegfbd adc ecfgbd cfdgb",
		"adg afgdc gcedfba cadfb bgfaed agbc cabdgf dgfec ag febacd": "fbecda cdgafb ga dga",
		"abcdfge egc egfabc ce gedabc dgfcb edgab gbdec fbaged adce": "ec gdbce ceg cbefga",
		"acbgf ebdfg gefadc fae efbag gcaefb ea cbea cgbafd cagdbfe": "bcae ea fdcgbae ecbfag",
	}
	return input
}
