# Pandora

## How to get Blowfish key?

~~~
jadx.bat com.pandora.android-21101001.apk
~~~

Result:

~~~java
com\pandora\constants\PandoraConstants.java
8:    public static final byte[] b = "6#26FRL$ZWD".getBytes(Charset.defaultCharset());
~~~

https://github.com/skylot/jadx

## How to get `password` for `partnerLogin`?

https://github.com/89z/googleplay/tree/master/cmd/mitmproxy-cert
