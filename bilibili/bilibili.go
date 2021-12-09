package bilibili

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "os"
    "os/exec"
    "path"
    "sort"
    "strconv"
    "time"
)

type Client struct {
    token     *string
    directory *string
}

func New(token string, directory string) Client {
    return Client{
        token:     &token,
        directory: &directory,
    }
}

var userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36"

type UserData struct {
    Code    float64 `json:"code"`
    Message string  `json:"message"`
    Ttl     float64 `json:"ttl"`
    Data    struct {
        //  IsLogin 是否已登录
        IsLogin bool `json:"isLogin"`
        //  EmailVerified 0:未验证 1:已验证
        EmailVerified uint8 `json:"email_verified"`
        //  Face 用户头像url
        Face string `json:"face"`
        // LevelInfo 等级信息
        LevelInfo struct {
            //  CurrentLevel 当前等级
            CurrentLevel uint64 `json:"current_level"`
            //  CurrentMin 当前等级经验最低值
            CurrentMin uint64 `json:"current_min"`
            //  CurrentExp 当前经验
            CurrentExp uint64 `json:"current_exp"`
            //  NextExp 升级下一等级需达到的经验
            NextExp uint64 `json:"next_exp"`
        } `json:"level_info"`
        //  Mid 用户UID
        Mid uint64 `json:"mid"`
        //  MobileVerified 是否验证手机号 0:未验证 1:已验证
        MobileVerified uint8 `json:"mobile_verified"`
        //  Money 拥有硬币数
        Money uint64 `json:"money"`
        //  Moral 当前节操值 上限70
        Moral uint64 `json:"moral"`
        //  Uname 用户昵称
        Uname string `jso:"uname"`
        //  VipDueDate 大会员到期时间	毫秒 时间戳
        VipDueDate uint64 `json:"vipDueDate"`
        //  VipStatus 会员开通状态 0:无 1:有
        VipStatus uint8 `json:"vipStatus"`
        //  VipType 大会员类型 0:无 1:月度 2:年度
        VipType uint8 `json:"vipType"`
        //  VipPayType 会员开通状态 0:无 1:有
        VipPayType uint8 `json:"vip_pay_type"`
        //  Wallet B币信息
        Wallet struct {
            //  Mid 登录用户UID
            Mid uint64 `json:"mid"`
            //  BcoinBalance 拥有B币数
            BcoinBalance uint64 `json:"bcoin_balance"`
            //  CouponBalance 每月奖励B币数
            CouponBalance uint64 `json:"coupon_balance"`
        }
    } `json:"data"`
}

func (c Client) GetUserData() (*UserData, error) {
    if c.token == nil {
        return nil, errors.New("SESSDATA is empty")
    }

    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://api.bilibili.com/nav", nil)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    req.Header.Set("user-agent", userAgent)
    req.Header.Set("cookie", fmt.Sprintf("SESSDATA=%v", *c.token))
    resp, err := client.Do(req)
    if resp == nil{
        // todo 编写错误信息
        return nil, errors.New("")
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        log.Println(resp.Status)
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    result := &UserData{}
    if err = json.Unmarshal(body, result); err != nil {
        return nil, err
    } else if result.Code != 0 {
        return nil, errors.New(result.Message)
    } else {
        return result, nil
    }
}

type VideoData struct {
    Code    float64 `json:"code"`
    Message string  `json:"message"`
    Ttl     float64 `json:"ttl"`
    Data    struct {
        //  Bvid 视频bvID
        Bvid string `json:"bvid"`
        //  Aid 视频avID
        Aid uint64 `json:"aid"`
        //  Videos 视频分P总数
        Videos uint64 `json:"videos"`
        //  Tid 分区ID
        Tid uint64 `json:"tid"`
        //  Tname 子分区名称
        Tname string `json:"tname"`
        //  Copyright 版权标志 1:自制 2:转载
        Copyright uint8 `json:"copyright"`
        //  Pic 视频封面图片url
        Pic string `json:"pic"`
        //  Title 视频标题
        Title string `json:"title"`
        //  Pubdate 视频上传时间
        Pubdate int64 `json:"pubdate"`
        //  Ctime 视频审核通过时间
        Ctime int64 `json:"ctime"`
        //  Desc 视频简介
        Desc string `json:"desc"`
        //  Duration 视频总计持续时长(所有分P) 单位:秒
        Duration uint64 `json:"duration"`
        //  Mission_id 视频参与的活动ID
        MissionId uint64 `json:"mission_id"`
        //  Owner 视频UP主信息
        Owner struct {
            //  Mid UP主UID
            Mid uint64 `json:"mid"`
            //  Name UP主昵称
            Name string `json:"name"`
            //  Face UP主头像
            Face string `json:"face"`
        } `json:"owner"`
        //  Stat 视频状态
        Stat struct {
            //  View 观看次数 屏蔽时 = -1
            View float64 `json:"view"`
            //  Danmaku 弹幕条数
            Danmaku uint64 `json:"danmaku"`
            //  Reply 评论条数
            Reply uint64 `json:"reply"`
            //  Favorite 收藏人数
            Favorite uint64 `json:"favorite"`
            //  Coin 投币枚数
            Coin uint64 `json:"coin"`
            //  Share 分享次数
            Share uint64 `json:"share"`
            //  Like 获赞次数
            Like uint64 `json:"like"`
        } `json:"stat"`
        //  Pages 视频分P列表
        Pages []PartItem `json:"pages"`
    } `json:"data"`
}

type PartItem struct {
    //  Cid 当前分P的CID
    Cid uint64 `json:"cid"`
    //  Page 当前分P
    Page uint64 `json:"page"`
    //  From 视频来源 vupload:普通上传 hunan:芒果TV
    From string `json:"from"`
    //  Part 当前分P标题
    Part string `json:"part"`
    //  Duration 当前分P持续时间 单位:秒
    Duration float64 `json:"duration"`
}

func (c Client) GetVideoData(bvid string) (*VideoData, error) {
    if c.token == nil {
        return nil, errors.New("SESSDATA is empty\n")
    }

    urlParams := url.Values{}
    urlParams.Set("bvid", bvid)

    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/view", nil)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    req.Header.Set("user-agent", userAgent)
    req.Header.Set("cookie", fmt.Sprintf("SESSDATA=%v", *c.token))
    req.URL.RawQuery = urlParams.Encode()

    resp, err := client.Do(req)
    if resp == nil{
        // todo 编写错误信息
        return nil, errors.New("")
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        log.Println(resp.Status)
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    result := &VideoData{}
    if err = json.Unmarshal(body, result); err != nil {
        return nil, err
    } else {
        return result, nil
    }
}

type StreamFormat struct {
    Quality     uint64 `json:"quality"`
    Format      string `json:"format"`
    Description string `json:"description"`
    DisplayDesc string `json:"display_desc"`
    Superscript string `json:"superscript"`
}

type FlvStream []struct {
    //  Order 视频分段序号 某些视频会分为多个片段（从1顺序增长）
    Order uint64 `json:"order"`
    //  Length 视频长度 单位:毫秒
    Length uint64 `json:"length"`
    //  Size   视频大小 单位:Byte
    Size uint64 `json:"size"`
    //  Url   视频流url url内容存在转义符
    Url string `json:"url"`
    //  BackupUrl 	备用视频流
    BackupUrl []string `json:"backup_url"`
}

func (f FlvStream) Len() int {
    return len(f)
}
func (f FlvStream) Less(i, j int) bool {
    return f[i].Order > f[j].Order
}
func (f FlvStream) Swap(i, j int) {
    f[i].Order, f[j].Order = f[j].Order, f[i].Order
}

type DashStreamVideo []struct {
    Id        uint64 `json:"id"`
    BaseUrl   string `json:"baseUrl"`
    MimeType  string `json:"mimeType"`
    Bandwidth uint64 `json:"bandwidth"`
    Codecs    string `json:"codecs"`
    Width     uint64 `json:"width"`
    Height    uint64 `json:"height"`
    FrameRate string `json:"frameRate"`
}

func (v DashStreamVideo) Len() int {
    return len(v)
}
func (v DashStreamVideo) Less(i, j int) bool {
    return v[i].Id > v[j].Id && v[i].Bandwidth > v[j].Bandwidth
}
func (v DashStreamVideo) Swap(i, j int) {
    v[i], v[j] = v[j], v[i]
}

type DashStreamAudio []struct {
    Id        uint64 `json:"id"`
    BaseUrl   string `json:"baseUrl"`
    Bandwidth uint64 `json:"bandwidth"`
    MimeType  string `json:"mimeType"`
    Codecs    string `json:"codecs"`
}

func (a DashStreamAudio) Len() int {
    return len(a)
}
func (a DashStreamAudio) Less(i, j int) bool {
    return a[i].Id > a[j].Id
}
func (a DashStreamAudio) Swap(i, j int) {
    a[i].Id, a[j].Id = a[j].Id, a[i].Id
}

type DashStream struct {
    Video DashStreamVideo `json:"video"`
    Audio DashStreamAudio `json:"audio"`
}

type StreamQuality []uint64

func (q StreamQuality) Len() int {
    return len(q)
}
func (q StreamQuality) Less(i, j int) bool {
    return q[i] > q[j]
}
func (q StreamQuality) Swap(i, j int) {
    q[i], q[j] = q[j], q[i]
}

type VideoStream struct {
    Code    float64 `json:"code"`
    Message string  `json:"message"`
    Ttl     float64 `json:"ttl"`
    Data    *struct {
        Quality           uint64         `json:"quality"`
        Format            string         `json:"format"`
        Timelength        uint64         `json:"timelength"`
        AcceptFormat      string         `json:"accept_format"`
        AcceptDescription []string       `json:"accept_description"`
        AcceptQuality     StreamQuality  `json:"accept_quality"`
        SupportFormats    []StreamFormat `json:"support_formats"`
        Durl              *FlvStream     `json:"durl"`
        Dash              *DashStream    `json:"dash"`
    } `json:"data"`
}

func (c Client) GetVideoQuality(bvid string, cid uint64) ([]uint64, error) {
    if c.token == nil {
        return nil, errors.New("SESSDATA is empty\n")
    }

    urlParams := url.Values{}
    urlParams.Set("bvid", bvid)
    urlParams.Set("cid", strconv.FormatUint(cid, 10))
    urlParams.Set("fourk", "1")

    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://api.bilibili.com/x/player/playurl", nil)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    req.Header.Set("user-agent", userAgent)
    req.Header.Set("cookie", fmt.Sprintf("SESSDATA=%v", *c.token))
    req.URL.RawQuery = urlParams.Encode()

    resp, err := client.Do(req)
    if resp == nil{
        // todo 编写错误信息
        return nil, errors.New("")
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        log.Println(resp.Status)
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    result := &VideoStream{}
    if err = json.Unmarshal(body, result); err != nil {
        return nil, err
    } else if result.Data != nil {
        sort.Sort(result.Data.AcceptQuality)

        return result.Data.AcceptQuality, nil
    } else {
        return nil, errors.New(result.Message)
    }
}

func (c Client) GetVideoStream(bvid string, cid uint64, qualityId uint64) (*VideoStream, error) {
    if c.token == nil {
        return nil, errors.New("SESSDATA is empty\n")
    }

    urlParams := url.Values{}
    urlParams.Set("bvid", bvid)
    urlParams.Set("cid", strconv.FormatUint(cid, 10))
    urlParams.Set("qn", strconv.FormatUint(qualityId, 10))
    urlParams.Set("fourk", "1")
    urlParams.Set("fnval", "16")

    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://api.bilibili.com/x/player/playurl", nil)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    req.Header.Set("user-agent", userAgent)
    req.Header.Set("cookie", fmt.Sprintf("SESSDATA=%v", *c.token))
    req.URL.RawQuery = urlParams.Encode()

    resp, err := client.Do(req)
    if resp == nil{
        // todo 编写错误信息
        return nil, errors.New("")
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        log.Println(resp.Status)
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    result := &VideoStream{}
    if err = json.Unmarshal(body, result); err != nil {
        return nil, err
    } else if result.Data != nil {
        if result.Data.Dash != nil {
            sort.Sort(result.Data.Dash.Video)
            sort.Sort(result.Data.Dash.Audio)
        }

        return result, nil
    } else {
        return nil, errors.New(result.Message)
    }
}

func (c Client) Download(bvid string, url string) (*string, error) {
    if c.token == nil {
        return nil, errors.New("SESSDATA is empty\n")
    }

    if c.directory == nil {
        return nil, errors.New("Directory is empty\n")
    }

    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    req.Header.Set("user-agent", userAgent)
    req.Header.Set("cookie", fmt.Sprintf("SESSDATA=%v", *c.token))
    req.Header.Set("referer", fmt.Sprintf("https://www.bilibili.com/video/%v", bvid))

    resp, err := client.Do(req)
    if resp == nil{
        // todo 编写错误信息
        return nil, errors.New("")
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        log.Println(resp.Status)
        return nil, err
    }

    total := resp.Header.Get("content-length")
    fileName := fmt.Sprintf("%v-%v.temp", total, strconv.FormatInt(time.Now().Unix(), 10))
    outFilePath := path.Join(*c.directory, fileName)

    isExists := c.pathExists(*c.directory)
    if !isExists {
        if err = os.MkdirAll(*c.directory, 0777); err != nil {
            return nil, err
        }
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    if err = ioutil.WriteFile(outFilePath, body, 777); err != nil {
        log.Println(err)
    }

    return &outFilePath, nil
}

func (c Client) pathExists(path string) bool {
    _, err := os.Stat(path)
    if err != nil {
        if os.IsExist(err) {
            return true
        }
        return false
    }
    return true
}

func (c Client) Merge(outFileName string, paths ...string) error {
    defer func() {
        for _, p := range paths {
            _ = os.Remove(p)
        }
    }()

    outPath := path.Join(*c.directory, fmt.Sprintf("%v.mkv", outFileName))
    arg := make([]string, 0)

    for _, p := range paths {
        arg = append(arg, "-i")
        arg = append(arg, p)
    }

    arg = append(arg, "-vcodec")
    arg = append(arg, "copy")
    arg = append(arg, "-acodec")
    arg = append(arg, "copy")
    arg = append(arg, outPath)

    cmd := exec.Command("ffmpeg", arg...)

    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr

    return cmd.Run()
}
