# kinhentinho

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