// Code generated by "pipeline"; DO NOT EDIT.
package diona

import (
	_ "embed"

	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
}

func (x *char) Data() *model.AvatarData {
	return base
}

var (
	auto = [][]float64{
		auto_1,
		auto_2,
		auto_3,
		auto_4,
		auto_5,
	}
)

var (
	// attack: aim = [6]
	aim = []float64{
		0.43860000371932983,
		0.47429999709129333,
		0.5099999904632568,
		0.5609999895095825,
		0.5967000126838684,
		0.637499988079071,
		0.6935999989509583,
		0.7497000098228455,
		0.8058000206947327,
		0.8669999837875366,
		0.937125027179718,
		1.019592046737671,
		1.102059006690979,
		1.184525966644287,
		1.2744899988174438,
	}
	// attack: auto_1 = [0]
	auto_1 = []float64{
		0.3612000048160553,
		0.3905999958515167,
		0.41999998688697815,
		0.4620000123977661,
		0.49140000343322754,
		0.5249999761581421,
		0.5712000131607056,
		0.6173999905586243,
		0.6636000275611877,
		0.7139999866485596,
		0.7717499732971191,
		0.8396639823913574,
		0.9075779914855957,
		0.975492000579834,
		1.0495799779891968,
	}
	// attack: auto_2 = [1]
	auto_2 = []float64{
		0.3353999853134155,
		0.3626999855041504,
		0.38999998569488525,
		0.42899999022483826,
		0.4562999904155731,
		0.48750001192092896,
		0.5303999781608582,
		0.5733000040054321,
		0.6161999702453613,
		0.6629999876022339,
		0.7166249752044678,
		0.7796880006790161,
		0.8427510261535645,
		0.905813992023468,
		0.9746099710464478,
	}
	// attack: auto_3 = [2]
	auto_3 = []float64{
		0.45579999685287476,
		0.492900013923645,
		0.5299999713897705,
		0.5830000042915344,
		0.6201000213623047,
		0.6625000238418579,
		0.72079998254776,
		0.7791000008583069,
		0.8374000191688538,
		0.9010000228881836,
		0.9738749861717224,
		1.0595760345458984,
		1.1452770233154297,
		1.230978012084961,
		1.324470043182373,
	}
	// attack: auto_4 = [3]
	auto_4 = []float64{
		0.4300000071525574,
		0.4650000035762787,
		0.5,
		0.550000011920929,
		0.5849999785423279,
		0.625,
		0.6800000071525574,
		0.7350000143051147,
		0.7900000214576721,
		0.8500000238418579,
		0.918749988079071,
		0.9995999932289124,
		1.0804500579833984,
		1.1612999439239502,
		1.249500036239624,
	}
	// attack: auto_5 = [4]
	auto_5 = []float64{
		0.5375000238418579,
		0.581250011920929,
		0.625,
		0.6875,
		0.731249988079071,
		0.78125,
		0.8500000238418579,
		0.918749988079071,
		0.987500011920929,
		1.0625,
		1.1484379768371582,
		1.249500036239624,
		1.3505619764328003,
		1.4516249895095825,
		1.5618749856948853,
	}
	// attack: fullaim = [7]
	fullaim = []float64{
		1.2400000095367432,
		1.3329999446868896,
		1.4259999990463257,
		1.5499999523162842,
		1.6430000066757202,
		1.7359999418258667,
		1.8600000143051147,
		1.9839999675750732,
		2.1080000400543213,
		2.2320001125335693,
		2.360960006713867,
		2.529599905014038,
		2.698240041732788,
		2.866879940032959,
		3.035520076751709,
	}
	// skill: paw = [0]
	paw = []float64{
		0.41920000314712524,
		0.4506399929523468,
		0.48208001255989075,
		0.5239999890327454,
		0.5554400086402893,
		0.5868800282478333,
		0.6287999749183655,
		0.6707199811935425,
		0.7126399874687195,
		0.7545599937438965,
		0.7964800000190735,
		0.8384000062942505,
		0.8907999992370605,
		0.9431999921798706,
		0.9955999851226807,
	}
	// skill: pawDur = [5]
	pawDur = []float64{
		1.7999999523162842,
		1.899999976158142,
		2,
		2.0999999046325684,
		2.200000047683716,
		2.299999952316284,
		2.4000000953674316,
		2.4000000953674316,
		2.4000000953674316,
		2.4000000953674316,
		2.4000000953674316,
		2.4000000953674316,
		2.4000000953674316,
		2.4000000953674316,
		2.4000000953674316,
	}
	// skill: pawShieldFlat = [2]
	pawShieldFlat = []float64{
		692.8065795898438,
		762.0968627929688,
		837.1613159179688,
		918,
		1004.6128540039062,
		1096.9998779296875,
		1195.1611328125,
		1299.0965576171875,
		1408.80615234375,
		1524.2900390625,
		1645.5479736328125,
		1772.5802001953125,
		1905.3865966796875,
		2043.9671630859375,
		2188.322021484375,
	}
	// skill: pawShieldPer = [1]
	pawShieldPer = []float64{
		0.07199999690055847,
		0.07739999890327454,
		0.0828000009059906,
		0.09000000357627869,
		0.09539999812841415,
		0.10080000013113022,
		0.1080000028014183,
		0.1151999980211258,
		0.12240000069141388,
		0.12960000336170197,
		0.13680000603199005,
		0.14399999380111694,
		0.15299999713897705,
		0.16200000047683716,
		0.17100000381469727,
	}
	// burst: burst = [0]
	burst = []float64{
		0.800000011920929,
		0.8600000143051147,
		0.9200000166893005,
		1,
		1.059999942779541,
		1.1200000047683716,
		1.2000000476837158,
		1.2799999713897705,
		1.3600000143051147,
		1.440000057220459,
		1.5199999809265137,
		1.600000023841858,
		1.7000000476837158,
		1.7999999523162842,
		1.899999976158142,
	}
	// burst: burstDot = [1]
	burstDot = []float64{
		0.5264000296592712,
		0.5658800005912781,
		0.6053599715232849,
		0.6579999923706055,
		0.6974800229072571,
		0.7369599938392639,
		0.7896000146865845,
		0.8422399759292603,
		0.8948799967765808,
		0.9475200176239014,
		1.0001599788665771,
		1.0528000593185425,
		1.1186000108718872,
		1.184399962425232,
		1.2502000331878662,
	}
	// burst: burstHealFlat = [3]
	burstHealFlat = []float64{
		513.1900634765625,
		564.5161743164062,
		620.1195068359375,
		680,
		744.1576538085938,
		812.592529296875,
		885.3045654296875,
		962.2937622070312,
		1043.5601806640625,
		1129.103759765625,
		1218.9244384765625,
		1313.0223388671875,
		1411.3974609375,
		1514.0498046875,
		1620.979248046875,
	}
	// burst: burstHealPer = [2]
	burstHealPer = []float64{
		0.053360000252723694,
		0.05736200138926506,
		0.06136399880051613,
		0.06669999659061432,
		0.07070200145244598,
		0.07470399886369705,
		0.08004000037908554,
		0.08537600189447403,
		0.09071200340986252,
		0.09604799747467041,
		0.1013839989900589,
		0.10672000050544739,
		0.1133899986743927,
		0.12005999684333801,
		0.12672999501228333,
	}
)