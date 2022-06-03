# kinhentinho

## Simply echoes the current request in curl-able form, thanks to [http2curl](https://github.com/moul/http2curl)

```
$ docker run -i --net=host public.ecr.aws/s8t0c2z4/kinhentinho:latest # ...then fire up ngrok http 8000 and browse it
curl -X 'GET' -d '' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-US,en;q=0.9,pt;q=0.8,es;q=0.7' -H 'Sec-Ch-Ua: " Not A;Brand";v="99", "Chromium";v="102", "Google Chrome";v="102"' -H 'Sec-Ch-Ua-Mobile: ?0' -H 'Sec-Ch-Ua-Platform: "Linux"' -H 'Sec-Fetch-Dest: document' -H 'Sec-Fetch-Mode: navigate' -H 'Sec-Fetch-Site: none' -H 'Sec-Fetch-User: ?1' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36' -H 'X-Forwarded-For: 2800:e2:8880:1454:963a:3c56:b867:3014' -H 'X-Forwarded-Proto: https' 'https://127.0.0.1:8000/'
curl -X 'GET' -d '' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-US,en;q=0.9' -H 'Referer: https://www.google.com/' -H 'Sec-Fetch-Dest: document' -H 'Sec-Fetch-Mode: navigate' -H 'Sec-Fetch-Site: none' -H 'Sec-Fetch-User: ?1' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36' -H 'X-Forwarded-For: 181.48.255.69' -H 'X-Forwarded-Proto: https' 'https://127.0.0.1:8000/'
curl -X 'GET' -d '' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-US,en;q=0.9,pt;q=0.8,es;q=0.7' -H 'Cache-Control: max-age=0' -H 'Sec-Ch-Ua: " Not A;Brand";v="99", "Chromium";v="102", "Google Chrome";v="102"' -H 'Sec-Ch-Ua-Mobile: ?0' -H 'Sec-Ch-Ua-Platform: "Linux"' -H 'Sec-Fetch-Dest: document' -H 'Sec-Fetch-Mode: navigate' -H 'Sec-Fetch-Site: none' -H 'Sec-Fetch-User: ?1' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36' -H 'X-Forwarded-For: 2800:e2:8880:1454:963a:3c56:b867:3014' -H 'X-Forwarded-Proto: https' 'https://127.0.0.1:8000/'
curl -X 'GET' -d '' -H 'Accept: image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-US,en;q=0.9,pt;q=0.8,es;q=0.7' -H 'Referer: https://551a-2800-e2-8880-1454-963a-3c56-b867-3014.ngrok.io/' -H 'Sec-Ch-Ua: " Not A;Brand";v="99", "Chromium";v="102", "Google Chrome";v="102"' -H 'Sec-Ch-Ua-Mobile: ?0' -H 'Sec-Ch-Ua-Platform: "Linux"' -H 'Sec-Fetch-Dest: image' -H 'Sec-Fetch-Mode: no-cors' -H 'Sec-Fetch-Site: same-origin' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36' -H 'X-Forwarded-For: 2800:e2:8880:1454:963a:3c56:b867:3014' -H 'X-Forwarded-Proto: https' 'https://127.0.0.1:8000/favicon.ico'
```

## Why?

`kubectl set-image deployment/which-is-failing public.ecr.aws/s8t0c2z4/kinhentinho:latest` and you'll get your request curl-lable. 

Why not?

## Why Kinhentinho?

Its '500', in portuguese diminutive form (-inho, as my Spanish-speaking friends will mock me - I'm being preemptive)

The 'k' is just a hint to where it was meant to be used

## Which port?

By default, it listens on `8000`. You can set the listen addr via the environment variable `LISTEN_ADDR`. 

e.g., to make it listen on 8080, make `LISTEN_ADDR` equal to `:8080`

# TODO

  * Make it smaller (with `scratch` image)