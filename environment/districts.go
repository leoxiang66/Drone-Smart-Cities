package environment

type District struct {
    Name string
}

// 全局变量
var DISTRICTS []District

func init() {
    // will be called initially and automatically
    districtNames := []string{
        "Central and Western",
        "Eastern",
        "Southern",
        "Wan Chai",
        "Kowloon City",
        "Kwun Tong",
        "Sham Shui Po",
        "Wong Tai Sin",
        "Yau Tsim Mong",
        "Islands",
        "Kwai Tsing",
        "North",
        "Sai Kung",
        "Sha Tin",
        "Tai Po",
        "Tsuen Wan",
        "Tuen Mun",
        "Yuen Long",
    }

    DISTRICTS = make([]District, len(districtNames))
    for i, name := range districtNames {
        DISTRICTS[i] = District{Name: name}
    }
}
