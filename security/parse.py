import re 
import pandas as pd

class Request:
    def __init__(self, ipaddr=None, timeinfo=None, requestinfo=None, method=None, path=None, protocol=None, status=None, size=None, referer=None, ua=None, cookie=None):
        self.ipaddr = ipaddr
        self.timeinfo = timeinfo
        self.method = method
        self.path = path
        self.protocol = protocol
        self.status = status
        self.size = size
        self.referer = referer
        self.ua = ua
        self.cookie = cookie
        self.requestinfo = requestinfo
    
    def to__dict(self):
        return {
            "ipaddr": self.ipaddr,
            "timeinfo": self.timeinfo,
            "method": self.method,
            "path": self.path,
            "protocol": self.protocol,
            "status": self.status,
            "size": self.size,
            "referer": self.referer,
            "ua": self.ua,
            "cookie": self.cookie,
            "requestinfo": self.requestinfo
        }

def parseLogFile(filename):
    requests = []
    with open(filename, 'r') as f:
        lines = f.readlines()
        for line in lines:
            try:
                r = parseLog(line).to__dict()
                requests.append(r)
            except Exception as e:
                print(e)
                print(line)
                print('------------------')

    df = pd.DataFrame(requests)
    # print(df.groupby('ipaddr').size().reset_index(name='counts').sort_values('counts', ascending=False))
    # print(df.groupby('path').size().reset_index(name='counts').sort_values('counts', ascending=False))
    print(df.groupby('ua').size().reset_index(name='counts').sort_values('counts', ascending=False))


def parseLog(s):
    regex = r'(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}) - - \[(.*)\] "(.*)" (\d{3}) (\d+) "(.*)" "(.*)" "(.*)"'
    matches = re.finditer(regex, s)
    requests = []
    for match in matches:
        ipaddr= match.group(1)
        timeinfo = match.group(2)
        request = match.group(3)
        if len(request.split(' ')) ==2:
            method, path = request.split(' ')
            protocol = None
            requestinfo = None
        elif len(request.split(' ')) ==3:
            method, path, protocol = request.split(' ')
            requestinfo = None
        elif len(request.split(' ')) ==1 or len(request.split(' ')) >3:
            requestinfo = request
            method = None
            path = None
            protocol = None
        status = match.group(4)
        size = match.group(5)
        referer = match.group(6)
        ua = match.group(7)
        cookie = match.group(8)
        r = Request(ipaddr, timeinfo,requestinfo, method, path, protocol, status, size, referer, ua, cookie)
        requests.append(r)
    return requests[0] if len(requests) > 0 else None


if __name__ == "__main__":
    # s2 = '49.36.188.169 - - [23/Dec/2023:11:18:54 +0000] "GET /admin/bond/bond/ HTTP/1.1" 200 21608 "https://fintrekkers.com/admin/bullion/bullionpricehistory/" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36" "-"'
    # s1='172.200.58.53 - - [23/Dec/2023:16:10:13 +0000] "GET //priv8.php HTTP/1.1" 404 179 "-" "Go-http-client/1.1" "-"'
    # s3 = '172.200.58.53 - - [23/Dec/2023:16:10:39 +0000] "GET //wp-content/plugins/wp-file-upload/ROOBOTS.php HTTP/1.1" 404 179 "-" "Go-http-client/1.1" "-"'
    # s4='118.26.37.23 - - [23/Dec/2023:16:54:56 +0000] "GET /h5/ HTTP/1.1" 503 599 "https://fintrekkers.com/h5/" "Mozilla/5.0 (Linux; Android 11; vivo 1906; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/87.0.4280.141 Mobile Safari/537.36 VivoBrowser/8.9.0.0" "-"'
    # s5='62.233.50.179 - - [23/Dec/2023:18:41:03 +0000] "\x03\x00\x00/*\xE0\x00\x00\x00\x00\x00Cookie: mstshash=Administr" 400 157 "-" "-" "-"'
    # s6='205.210.31.47 - - [23/Dec/2023:18:57:47 +0000] "\x16\x03\x01\x00\xEE\x01\x00\x00\xEA\x03\x03Z\xF7Wd\xC2^\x92\x9F\xFD\x92\x99\xE3zmI\x1DJ|\xC4r\xC2=\x01\x09\x5CuN\x83\xF21\xCA\xF0 \xA9jdR\xDE\xE9R9Q\xB1\x95n\x02M\xAA\xFFDW\xB2\xCE\xB4\xA7\x96\xF9R>h\xDA#\xC7p/\x00&\xC0+\xC0/\xC0,\xC00\xCC\xA9\xCC\xA8\xC0\x09\xC0\x13\xC0" 400 157 "-" "-" "-"'
    # parseLog(s1)
    # parseLog(s2)
    # parseLog(s3)
    # parseLog(s4)
    # parseLog(s5)
    # parseLog(s6)
    parseLogFile('hackerlogs.log')