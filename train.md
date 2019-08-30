SSH                    BSD General Commands Manual                   SSH



NAME

    ssh -- OpenSSH SSH 客户端工具(远程登录程序)



SYNOPSIS

    ssh [-1246AaCfGgKkMNnqsTtVvXxYy] [-b bind_address] [-c cipher_spec]

        [-D [bind_address:]port] [-E log_file] [-e escape_char]

        [-F configfile] [-I pkcs11] [-i identity_file] [-L address]

        [-l login_name] [-m mac_spec] [-O ctl_cmd] [-o option] [-p port]

        [-Q query_option] [-R address] [-S ctl_path] [-W host:port]

        [-w local_tun[:remote_tun]] [user@]hostname [command]



DESCRIPTION

    ssh(SSH客户端)是一个登陆远程主机和在远程主机上执行命令的程序。它的目

    的是在不安全的网络中为两个互不信任的主机提供安全加密的通信方式。也

    可以通过安全隧道被转发X11连接、任意TCP端口和UNIX套接字上的数据包。

    ssh连接并登录指定的主机(还可以指定用户名)。客户端必须提供身份标识给

    远程主机，提供方式有多种，见下文。

    如果ssh命令行中指定了命令，则将在远程主机上执行而不是登录远程主机。



    选项说明如下：

    -1      强制使用ssh v1版本。

    -2      强制使用ssh v2版本。

    -4      强制只使用IPv4地址。

    -6      强制只使用IPv6地址。

    -A      启用代理转发功能，也可在全局配置文件(/etc/ssh/ssh_config)中配置。

            代理转发功能应该要谨慎开启。

    -a      禁用代理转发功能。

    -b bind_address

            在本地主机上绑定用于ssh连接的地址，当系统有多个ip时才生效。

    -C      请求会话间的数据压缩传递。对于网络缓慢的主机，压缩对连接有所

            提升。但对网络流畅的主机来说，压缩只会更糟糕。

    -c      选择ssh会话间数据加密算法。

    -D [bind_address:]port

            指定一个本地动态应用层端口做转发端口。工作方式是分配一个套接
            字监听在此端口，当监听到此端口有连接时，此连接中的数据将通过
            安全隧道转发到server端，server端再和目的地(端口)建立连接，目
            的地(端口)由应用层协议决定。目前支SOCK4和SOCK5两种协议，并且
            SSH将扮演SOCKS服务端角色。

            只有root用户可以开启特权端口。动态转发端口也可以在配置文件
            中指定。

            默认情况下，转发端口将绑定在GatewayPorts指令指定的地址上，但
            是可以显式指定bind_address，如果bind_address设置为'localhost'，

            则转发端口将绑定在回环地址上，如果bind_address不设置或设置为
            '*'，则转发端口绑定在所有网路接口上。

    -E log_file

            将debug日志写入到log_file中，而不是默认的标准错误输出stderr。

    -e escape_char

            设置逃逸首字符，默认为'~'，设置为'none'将禁用逃逸字符，并使
            得会话完全透明。详细用法见后文。

    -F configfile

            指定用户配置文件，默认为~/.ssh/config，如果在命令行指定了该
            选项，则全局配置文件/etc/ssh_config将被忽略。

    -f      请求ssh在工作在后台模式。该选项隐含了'-n'选项，所以标准输入
            将变为/dev/null。

    -G      使用该选项将使得ssh在匹配完Host后将输出与之对应的配置选项，
            然后退出

    -g      允许远程主机连接到本地转发端口上。

    -I pkcs11

            Specify the PKCS#11 shared library ssh should use to communicate

            with a PKCS#11 token providing the user's private RSA key.

    -i identity_file

            指定公钥认证时要读取的私钥文件。默认为~/.ssh/id_rsa。

    -K      启用GSSAPI认证并将GSSAPI凭据转发(分派)到服务端。

    -k      禁止转发(分派)GSSAPI凭据到服务端。

    -L [bind_address:]port:host:hostport

    -L [bind_address:]port:remote_socket

    -L local_socket:host:hostport

    -L local_socket:remote_socket

            对本地指定的TCP端口port的连接都将转发到指定的远程主机及其端
            口上(host:hostport)。工作方式是在本地端分配一个socket监听TCP
            端口。当监听到本地此端口有连接时，连接将通过安全隧道转发给
            远程主机(server)，然后从远程主机(是server端)上建立一个到
            host:hostport的连接，完成数据转发。

            译者注：隧道建立在本地和远程主机(server端，即中间主机)之间，
            而非本地和host之间，也不是远程主机和host之间。

            端口转发也可以在配置文件中指定。只有root用户才能转发特权端口
            (小于1024)。

            默认本地端口被绑定在GatewayPorts指令指定的地址上。但是，显式
            指定的bind_address可以用于绑定连接到指定的地址上。如果设置

            bind_address为'localhost'，则表示被绑定的监听端口只可用于本地
            连接(即该端口监听在回环地址上)，如果不设置bind_address或设置
            为'*'则表示绑定的端口可用于所有网络接口上的连接(即表示该端口
            监听在所有地址上)。

    -l login_name

            指定登录在远程机器上的用户名。这也可以在全局配置文件中设置。

    -M      将ssh客户端置入'master'模式，以便连接共享(连接复用)。
            即实现ControlMaster和ControlPersist的相关功能。

    -m mac_spec

            A comma-separated list of MAC (message authentication code)
            algorithms, specified in order of preference. See the MACs
            key‐word for more information.

    -N      明确表示不执行远程命令。仅作端口转发时比较有用。

    -n      将/dev/null作为标准输入stdin，可以防止从标准输入中读取内容。
            当ssh在后台运行时必须使用该项。但当ssh被询问输入密码时失效。

    -O ctl_cmd

            Control an active connection multiplexing master process.  When
            the -O option is specified, the ctl_cmd argument is interpreted
            and passed to the master process.  Valid commands are: 'check'
            (check that the master process is running), 'forward' (request
            forwardings without command execution), 'cancel' (cancel for‐
            wardings), 'exit' (request the master to exit), and 'stop'
            (request the master to stop accepting further multiplexing
            requests).

    -o option

            Can be used to give options in the format used in the configura‐
            tion file.  This is useful for specifying options for which there
            is no separate command-line flag.  For full details of the
            options listed below, and their possible values, see

            ssh_config.

                    AddKeysToAgent

                    AddressFamily

                        连接时使用的地址系列。有效参数是'any'，'inet'（仅用于IPv4）
                        或'inet6'（仅用于IPv6）。

                    BatchMode

                        如果设置为'yes'，将禁用passphrase/password查询。
                        此选项在没有用户提供密码的脚本和其他批处理作业中很有用。
                        参数必须是'yes'或'no'。默认值为'no'。

                    BindAddress

                        使用本地计算机上指定的地址作为连接的源地址。
                        仅适用于具有多个地址的系统。请注意，
                        如果usePrivilegedPort设置为'yes'，则此选项不起作用。

                    CanonicalDomains

                    CanonicalizeFallbackLocal

                    CanonicalizeHostname

                    CanonicalizeMaxDots

                    CanonicalizePermittedCNAMEs

                    CertificateFile

                    ChallengeResponseAuthentication

                        指定是否使用challenge-response身份验证。
                        此关键字的参数必须是'yes'或'no'。默认值为'yes'。

                    CheckHostIP

                        如果此标志设置为'yes'，ssh将另外检查已知主机文件中的主机IP地址。
                        这允许ssh检测主机密钥是否因DNS欺骗而更改。
                        如果选项设置为'no'，则不会执行检查。默认值为'yes'。

                    Cipher

                        指定用于加密协议版本1中的会话的密码。
                        目前，支持'blowfish'，'3des'和'des'。
                        des仅在ssh 客户端中受支持，
                        以实现与不支持3des密码的传统协议1实现的互操作性。
                        由于加密弱点，强烈建议不要使用它。默认为'3des'。

                    Ciphers

                        按优先顺序指定协议版本2允许的密码。多个密码必须以逗号分隔。
                        支持的密码是 '3des-cbc'，'aes128-cbc'，'aes192-cbc'，
                        'aes256-cbc'，'aes128-ctr'，'aes192-ctr' '，'aes256-ctr'，
                        'arcfour128'，'arcfour256'，'arcfour'，'blowfish-cbc'
                        和'cast128-cbc'。

                    ClearAllForwardings

                        清除配置文件或命令行中指定的转发（本地、远程和动态端口）。
                        此选项主要作用是在使用ssh命令行时清除配置文件中的端口转发，
                        并由scp和sftp自动设置。参数必须是'yes'或'no'。
                        默认值为'no'。

                    Compression

                        指定是否使用压缩。参数必须是'yes'或'no'。默认值为'no'。

                    CompressionLevel

                        使用选项'Compression'时指定的'level'，
                        参数必须是从1（快速）到9（慢速，最佳）的整数。
                        默认level为6，
                        适用于大多数应用程序。这些值的含义与gzip 中的相同。
                        请注意，此选项仅适用于协议版本1。

                    ConnectionAttempts

                        指定退出前要进行的尝试次数（每秒一次）。参数必须是整数。
                        如果连接有时失败，这在脚本中可能很有用。默认值为1。

                    ConnectTimeout

                        指定连接到ssh服务器时使用的超时（秒），而不是使用默认系统tcp超时。
                        此值仅在目标关闭或确实无法访问时使用，而不是在它拒绝连接时使用。

                    ControlMaster

                        许在单个网络连接上共享多个会话。

                        当设置为'yes'时，ssh将侦听使用 controlPath 参数指定的控制套接字上的连接。
                        其他会话可以使用 ControlMaster 设置为 'no' 的同一 ControlPath 连接到此套接字（默认）。
                        这些会话将尝试重用主实例的网络连接，而不是启动新的网络连接，
                        但如果控制套接字不存在或未在侦听，则会恢复正常连接。
                        将此设置为'ask'将导致ssh侦听控制连接，
                        但需要在接受之前使用 SSH_ASKPASS 程序进行确认（有关详细信息，请参阅see  ssh-add）。
                        如果无法打开 controlPath，ssh将继续而不连接到主实例。
                        这些多路连接支持x11和 ssh-agent 转发，但转发的显示和代理将属于主连接，
                        即不能转发多个显示或代理。

                        另外两个选项允许机会多路复用：尝试使用主连接，但如果尚不存在，则回退到创建新连接。
                        这些选项是：'auto'和'autoask'。后者需要确认，如'ask'选项。

                    ControlPath

                        指定用于连接共享的控制套接字的路径。

                        如上面的ControlMaster部分所述，或者字符串'none'以禁用连接共享。
                        在路径中，'％l'将被本地主机名替换，'％h'将被目标主机名替换，
                        '％p'替换为端口，'％r'将被远程登录用户名替换。
                        建议用于机会连接共享的任何ControlPath至少包括％h，％p和％r。
                        这可确保唯一标识共享连接。

                    ControlPersist

                    DynamicForward

                        指定通过安全通道转发本地计算机上的TCP端口

                        然后使用应用程序协议确定从远程计算机连接到的位置。
                        参数必须是[bind_address:] port。
                        可以通过将地址括在方括号中或使用替代语法来指定IPv6地址：[bind_address/] port。

                        默认情况下，本地端口根据GatewayPorts设置进行绑定。
                        但是，可以使用显式bind_address将连接绑定到特定地址。
                        'localhost' 的bind_address表示侦听端口仅限本地使用，
                        而空地址或'*'表示该端口应该可从所有接口使用。
                        目前支持SOCKS4和SOCKS5协议，ssh将充当SOCKS服务器。
                        可以指定多个转发，并且可以在命令行上提供其他转发。
                        只有超级用户才能转发特权端口。

                    EnableSSHKeysign

                        在全局客户端配置文件 /etc/ssh/ssh_config 中将此选项设置为'yes'

                        可以在 HostbasedAuthentication 期间使用帮助程序 ssh-keysign。
                        参数必须是'yes'或'no'。默认为'no'。此选项应放在非特定于主机的部分中。

                        有关更多信息，请参阅 ssh-keysign。

                    EscapeChar

                        设置转义符（默认：'~'。转义符也可以在命令行上设置。

                        参数应为单个字符，'^'后跟字母或'none'，
                        以完全禁用转义字符（使连接对二进制数据透明）。

                    ExitOnForwardFailure

                        是否应终止连接，当指定的ssh无法设置所有动态请求，隧道端口、本地端口和远程端口转发。

                        参数必须是'yes'或'no'。默认值为'no'。

                    FingerprintHash

                    ForwardAgent

                        否将身份验证代理（如果有）转发到远程机器。

                        参数必须是'yes'或'no'。默认为'no'。应该慎重的使用代理转发。
                        能够绕过远程主机（使用了 Unix-domain socket 的代理）
                        权限的用户可以通过转发连接访问本地代理。虽然攻击者无法从代理获取密钥信息，
                        但是他们可以对密钥执行操作， 使其能够使用加载到代理中的身份得到身份验证。

                    ForwardX11

                        是否通过 secure channel 和 DISPLAY 自动重定向 x11 连接。

                        参数必须是'yes'或'no'。默认值为'no'。应谨慎启用X11转发。
                        能够绕过远程主机（对于用户的X11授权数据库）
                        的文件权限的用户可以通过转发的连接访问本地X11显示器。
                        如果还启用了ForwardX11Trusted选项，则攻击者可以执行诸如击键监视之类的活动。

                    ForwardX11Timeout

                    ForwardX11Trusted

                        如果此选项设置为'yes'，则远程X11客户端将具有对原始X11显示器的完全访问权限。

                        如果此选项设置为'no'，则远程X11客户端将被视为不受信任，
                        并且无法窃取或篡改属于受信任X11客户端的数据。

                        此外，用于会话的xauth将设置为在20分钟后过期。
                        在此之后，远程客户端将被拒绝访问。默认为'no'。

                        有关对不受信任的客户端施加的限制的完整详细信息，请参阅X11 SECURITY扩展规范。

                    GatewayPorts

                        是否允许远程主机连接到本地转发端口。

                        缺省情况下，ssh将本地端口转发绑定到 loopback address。
                        这可以防止其他远程主机连接到转发端口。
                        GatewayPorts可用于指定ssh应将本地端口转发绑定到通配符地址，
                        从而允许远程主机连接到转发端口。参数必须是'yes'或'no'。默认为'no'。

                    GlobalKnownHostsFile

                        用于全局主机密钥数据库的文件，而不是/etc/ssh/ssh_known_hosts。

                    GSSAPIAuthentication

                        是否允许基于GSSAPI的用户身份验证。默认为'no'。请注意，此选项仅适用于协议版本2。

                    GSSAPIKeyExchange

                        是否可以使用基于GSSAPI的密钥交换。

                        使用GSSAPI密钥交换时，服务器不需要主机密钥。默认为'否'。
                        请注意，此选项仅适用于协议版本2。

                    GSSAPIClientIdentity

                        如果是'set'，则指定ssh在连接到服务器时应使用的GSSAPI客户端标识。

                        默认设置为'unset'，这意味着将使用默认标识。

                    GSSAPIDelegateCredentials

                        转发（委托）证书到服务器。默认值为'否'。

                        请注意，此选项适用于使用GSSAPI的协议版本2连接。

                    GSSAPIRenewalForcesRekey

                        如果设置为'yes'，当强制更新ssh连接的密钥时更新客户端的证书，

                        使用兼容服务器可以委托连接上服务器的会话更新证书。默认值为'no'。

                    GSSAPITrustDns

                        设置为'yes'表示可信任的DNS以安全规范化地所连接主机的名称。

                        如果'no'，则在命令行中输入的主机名将原封不动地传递给GSSAPI库。

                        默认值为'no'。此选项仅适用于使用GSSAPI的协议版本2连接。

                    HashKnownHosts

                        当将host名和地址加入到 '~/.ssh/known_hosts' 中时，需要进行哈希加密。

                        这些被哈希加密的host一般是可以由ssh 和sshd使用的。
                        即使文件内容公开识别信息也不会被泄漏。默认值为'no'。
                        请注意，已知主机文件中的现有名称和地址不会自动转换，
                        但可以使用ssh-keygen手动哈希加密。

                    Host

                    HostbasedAuthentication

                        是否使用公钥身份验证尝试基于Rhosts的身份验证。

                        参数必须是'yes'或'no'。默认值为'no'。

                        此选项仅适用于协议版本2，与RhostsrsaAuthentication类似。

                    HostbasedKeyTypes

                    HostKeyAlgorithms

                        客户端希望按优先顺序使用的协议版本2主机密钥算法。

                        此选项的默认值为：'ssh-rsa，ssh-dss'。

                    HostKeyAlias

                        在主机密钥数据库文件中查找或保存主机密钥时，应使用的别名而不是实际主机名。

                        此选项对于隧道连接SSH连接或在单个主机上运行的多个服务器非常有用。

                    HostName

                        要登录的实际主机名。

                        这可用于指定主机的昵称或缩写。默认值是命令行上给定的名称。
                        也允许使用数字IP地址（只要符合命令行和主机名规范）。

                    IdentityFile

                        从中读取用户的RSA或DSA身份验证标识的文件。

                        协议版本1的默认值为 ~/.ssh/identity，
                        协议版本2的默认值为 ~/.ssh/id_rsa 和 ~/.ssh/id_dsa 。

                        此外，身份验证代理所代表的任何身份都将用于身份验证。
                        文件名可以使用波浪号语法来引用用户的主目录或以下转义字符之一：
                        '％d'（本地用户的主目录），'％u'（本地用户名），'％l'（本地 主机名），
                        '％h'（远程主机名）或'％r'（远程用户名）。
                        可以在配置文件中指定多个标识文件；所有这些标识都将按顺序尝试。

                    IdentitiesOnly

                        ssh 只能使用 ssh_config 文件中配置的身份验证文件，即使 ssh-agent 提供更多身份。

                        此关键字的参数必须为 'yes' 或 'no'。 此选项适用于 ssh-agent 提供许多不同身份的情况。 默认为 'no'

                    IPQoS

                    KbdInteractiveAuthentication

                        是否使用 keyboard-interactive 方式作为身份验证。

                        此关键字的参数必须为 'yes' 或 'no'； 默认为 'yes'

                    KbdInteractiveDevices

                        使用 keyboard-interactive 作为身份验证时的方法列表。

                        多个方法名称必须以逗号分隔。 默认是使用服务器指定的列表。
                        可用的方法取决于服务器支持的内容。
                        对于OpenSSH服务器，它可以是零或更多：'bsdauth'，'pam' 和 'skey'。

                    KexAlgorithms

                    LocalCommand

                        成功连接到服务器后在本地计算机上执行的命令。

                        命令字符串扩展到行的末尾，并使用用户的shell执行。
                        将执行以下转义字符替换：'％d'（本地用户的主目录），'％h'（远程主机名），
                        '％l'（本地主机名），'％n'（提供的主机名） 命令行），'％p'（远程端口），
                        '％r'（远程用户名）或'％u'（本地用户名）。

                        除非已启用PermitLocalCommand，否则将忽略此伪指令。

                    LocalForward

                        本地计算机上的TCP端口通过安全通道转发到远程计算机的指定主机和端口。

                        第一个参数必须是 [bind_address：] port，第二个参数必须是 host：hostport。
                        可以通过将地址括在方括号中或使用替代语法来指定IPv6地址：[bind_address /] port
                        和 host / hostport。 可以指定多个转发，并且可以在命令行上提供其他转发。
                        只有超级用户才能转发特权端口。 默认情况下，本地端口根据 GatewayPorts 设置进行绑定。
                        但是，可以使用显式 bind_address 将连接绑定到特定地址。
                        'localhost'的bind_address表示侦端口仅限本地使用，
                        而空地址或'*'表示该端口应该可从所有接口使用。

                    LogLevel

                        给出从 ssh 记录消息时使用的详细级别。

                        可能的值包括：QUIET，FATAL，ERROR，INFO，VERBOSE，DEBUG，DEBUG1，DEBUG2和DEBUG3。

                        默认值为INFO。 DEBUG和DEBUG1是等价的。 DEBUG2和DEBUG3各自指定更高级别的详细输出。

                    MACs

                        按优先顺序排列的MAC（消息认证代码）算法。

                        MAC算法在协议版本2中用于数据完整性保护。多个算法必须以逗号分隔。
                        默认值：hmac-md5,hmac-sha1,umac-64@openssh.com,
                        hmac-ripemd160,hmac-sha1-96,hmac-md5-96

                    Match

                    NoHostAuthenticationForLocalhost

                        如果主目录跨计算机共享，则可以使用此选项。

                        localhost将在每台机器上引用不同的机器，用户将收到许多关于更改的主机密钥的警告。

                        但是此选项会禁用localhost的主机身份验证。此关键字的参数必须是'yes'或'no'。
                        默认情况'no'，检查本地主机的主机密钥。

                    NumberOfPasswordPrompts

                        在放弃前，密码提示的次数。 此关键字的参数必须是整数。 默认值为3。

                    PasswordAuthentication

                        是否使用密码验证。 此关键字的参数必须为 'yes' 或 'no'。 默认为 'yes'。

                    PermitLocalCommand

                        允许通过LocalCommand选项执行本地命令，或使用ssh 中的 '!' 命令转义序列。
                        参数必须是 'yes' 或 'no'。 默认为 'no'。

                    PKCS11Provider

                    Port

                        要在远程主机上连接的端口号。 默认值为22。

                    PreferredAuthentications

                        客户端尝试协议2身份验证方法的顺序。

                        这允许客户端首选一种方法（e.g. keyboard-interactive）而不是另一种方法（如密码）。
                        此选项的默认值为：'gssapi-with-mic, hostbased, publickey,
                        keyboard-interactive, password'

                    Protocol

                        协议版本 ssh 应按优先顺序支持。可能的值为 '1' 和 '2'。

                        多个版本必须以逗号分隔。默认值为 '2,1'。
                        这意味着ssh尝试版本2，如果版本2不可用，则返回到版本1。

                    ProxyCommand

                        用于连接到服务器的命令。

                        命令字符串扩展到行的末尾，并使用用户的shell执行。
                        在命令字符串中，'%h' 将由要连接的主机名替换，'%p'将由端口替换。
                        该命令基本上可以是任何东西，应该从它的标准输入读取并写入到它的标准输出。
                        它最终应该连接在某台机器上运行的 sshd 服务器，或者在某处执行 sshd-i。
                        主机密钥管理将使用所连接主机的主机名进行（默认为用户键入的名称）。
                        将命令设置为 'none' 将完全禁用此选项。
                        请注意，CheckHostIP 不可用于使用代理命令进行连接。
                        该指令与nc 及其代理支持结合使用非常有用。

                        例如，以下指令将通过192.0.2.0的HTTP代理连接：
                        ProxyCommand /usr/bin/nc -X connect -x 192.0.2.0:8080 %h %p

                    ProxyUseFdpass

                    PubkeyAcceptedKeyTypes

                        是否尝试公钥身份验证。

                        此关键字的参数必须是 'yes' 或 'no'。默认值为 'yes'。此选项仅适用于协议版本2。

                    PubkeyAuthentication

                    RekeyLimit

                        在重新协商会话密钥之前可以传输的最大数据量。

                        参数是字节数，可选后缀为 'K'、'M' 或 'G'，
                        分别表示千字节、兆字节或千兆字节。
                        默认值介于 '1G' 和 '4G' 之间，具体取决于密码。此选项仅适用于协议版本2。

                    RemoteForward

                        远程计算机上的TCP端口通过安全通道转发到本地计算机的指定主机和端口。

                        第一个参数必须是[bind_address：] port，第二个参数必须是host：hostport。
                        可以通过将地址括在方括号中或使用替代语法来指定IPv6地址：
                        [bind_address /] port和host / hostport。
                        可以指定多个转发，并且可以在命令行上提供其他转发。
                        只有在远程计算机上以root用户身份登录时，才能转发特权端口。
                        如果port参数为'0'，则侦听端口将在服务器上动态分配，并在运行时报告给客户端。
                        如果未指定bind_address，则默认仅绑定到环回地址。
                        如果bind_address是'*'或空字符串，则请求转发侦听所有接口。
                        仅当启用了服务器的GatewayPorts选项时，才会指定远程bind_address（请参阅sshd_config）。

                    RequestTTY

                    RhostsRSAAuthentication

                        是否使用RSA主机身份验证尝试基于Rhosts的身份验证。

                        参数必须是 'yes' 或 'no'。默认值为 'no'。
                        此选项仅适用于协议版本1，并要求 ssh 为setuid root。

                    RSAAuthentication

                        是否尝试RSA身份验证。此关键字的参数必须是 'yes' 或 'no'。

                        只有在身份文件存在或身份验证代理正在运行时，才会尝试RSA身份验证。
                        默认值为 'no'。请注意，此选项仅适用于协议版本1。

                    SendEnv

                        应该将哪些来自本地的环境变量发送到服务器。

                        注意，只有协议2才支持环境传递。
                        服务器还必须支持它，并且必须将服务器配置为接受这些环境变量。
                        有关如何配置服务器的信息，请参阅sshd_config中的acceptenv。
                        变量由名称指定，名称中可能包含通配符。
                        多个环境变量可以用空格分隔，也可以跨多个 SendEnv 指令分布。

                        默认情况下不发送任何环境变量。

                    ServerAliveCountMax

                        设置服务器活动消息的数量（见下文），可以在没有ssh 从服务器接收任何消息的情况下发送。

                        如果在发送服务器活动消息时达到此阈值，则ssh将断开与服务器的连接，从而终止会话。
                        值得注意的是，服务器活动消息的使用与TCPKeepAlive（下面）非常不同。
                        服务器活动消息通过加密通道发送，因此不会是伪造的。

                        TCPKeepAlive启用的TCP keepalive选项是可伪造的。
                        当客户端或服务器依赖于知道连接何时变为非活动状态时，服务器活动机制很有价值。
                        默认值为3。

                        例如，如果ServerAliveInterval设置为15，并且ServerAliveCountMax保留为默认值，
                        则如果服务器没有响应，则ssh将在大约45秒后断开连接。此选项仅适用于协议版本2。

                    ServerAliveInterval

                        设置超时间隔（秒），在此间隔之后，如果没有从服务器接收到数据，
                        ssh 将通过加密通道发送消息以请求服务器的响应。

                        默认值为0，表示这些消息不会发送到服务器。此选项仅适用于协议版本2。

                    SmartcardDevice

                        使用哪种 Smartcard 设备。

                        此关键字的参数是设备 ssh 应该用于与用于存储用户的私有RSA密钥的 Smartcard 进行通信。

                        默认情况下，未指定任何设备，并且未激活智能卡支持。

                    StreamLocalBindMask

                    StreamLocalBindUnlink

                    StrictHostKeyChecking

                        如果此标志设置为 'yes'，则ssh将永远不会自动将主机密钥添加到 ~/.ssh/known_hosts文件，
                        并拒绝连接到主机密钥已更改的主机。

                        这提供了针对特洛伊木马攻击的最大保护，
                        但是当 /etc/ssh/ssh_known_hosts 文件维护不当或经常建立与新主机的连接时，这可能很烦人。
                        此选项强制用户手动添加所有新主机。
                        如果此标志设置为 'no'，则ssh将自动将新主机密钥添加到用户已知主机文件。
                        如果此标志设置为 'ask'，则只有在用户确认他们真正想要做的事情后，
                        才会将新的主机密钥添加到用户已知的主机文件中，
                        并且ssh拒绝连接到主机密钥已更改的主机。 在所有情况下，将自动验证已知主机的主机密钥。

                        参数必须是 'yes'，'no' 或 'ask'。 默认为 'ask'。

                    TCPKeepAlive

                        系统是否应向另一端发送 TCP keepalive 消息。

                        如果它们被发送，其中一台机器的连接或崩溃的死亡将被正确地注意到。
                        然而，这意味着如果线路暂时中断，连接就会中断，有些人觉得这很烦人。

                        默认值是 'yes'（发送 TCP keepalive消息），客户机将注意到网络是否关闭或远程主机是否已关闭。
                        这在脚本中很重要，许多用户也希望如此。

                        若要禁用 TCP Keepalive 消息，该值应设置为 'no'。

                    Tunnel

                        请求在客户端和服务器之间转发tun设备。

                        参数必须是 'yes', 'point-to-point' (layer 3), 'ethernet' (layer 2), or 'no'。
                        规定 'yes' 作为隧道模式的默认请求，即'点到点'。默认值为 'no'。

                    TunnelDevice

                        要在客户端（local_tun）和服务器（remote_tun）上打开的tun设备。

                        参数必须是local_tun [：remote_tun]。
                        设备可以通过数字ID或关键字 'any' 来指定，该关键字使用下一个可用的隧道设备。
                        如果未指定 remote_tun，则默认为 'any'。 默认为'any：any'。

                    UpdateHostKeys

                    UsePrivilegedPort

                        是否使用特权端口进行对外连接。参数必须是 'yes' 或 'no'。默认值为 'no'。

                        如果设置为 'yes'，ssh 必须为setuid root。
                        请注意，对于使用旧服务器的 RhostsrsaAuthentication，此选项必须设置为 'yes'。

                    User

                        作为登录的用户。当在不同的机器上使用不同的用户名时，这可能很有用。
                        这样就省去了在命令行中必须记住提供用户名的麻烦。

                    UserKnownHostsFile

                        用来代替 ~/.ssh/known_hosts 作为用户主机密钥数据库。

                    VerifyHostKeyDNS

                        指定是否使用DNS和sshfp资源记录验证远程密钥。

                        如果此选项设置为 'yes'，则客户端将绝对的信任来自 DNS 的安全密钥。
                        如果此选项设置为 'ask，不安全的密钥需要被处理（
                        显示密钥匹配的信息，需要根据 StrictHostKeyChecking 选项确认新的主机密钥）。

                    VisualHostKey

                        如果此标志设置为 'yes'，则除了登录时的十六进制密钥字符串和未知主机密钥外，
                        还会打印远程主机密钥的ASCII艺术表示。
                        如果此标志设置为 'no'，则登录时不会打印密钥字符串，
                        只会打印未知主机密钥的十六进制密钥字符串。 默认为 'no'。

                    XAuthLocation

                        xauth 程序的完整路径名。默认值为 /usr/bin/xauth。

     -p port

             指定要连接远程主机上哪个端口，也可在全局配置文件中指定。

     -Q query_option

             Queries ssh for the algorithms supported for the specified ver‐

             sion 2.  The available features are: cipher (supported symmetric

             ciphers), cipher-auth (supported symmetric ciphers that support

             authenticated encryption), mac (supported message integrity

             codes), kex (key exchange algorithms), key (key types), key-cert

             (certificate key types), key-plain (non-certificate key types),

             and protocol-version (supported SSH protocol versions).

     -q      静默模式。大多数警告信息将不输出。

     -R [bind_address:]port:host:hostport

     -R [bind_address:]port:local_socket

     -R remote_socket:host:hostport

     -R remote_socket:local_socket

             对远程(server端)指定的TCP端口port的连接都就将转发到本地主机和

             端口上，工作方式是在远端(server)分配一个套接字socket监听TCP端

             口。当监听到此端口有连接时，连接将通过安全隧道转发给本地，然后

             从本地主机建一条到host:hostport的连接。

             端口转发也可以在配置文件中指定。只有root用户才能转发特权端口

             (小于1024)。

             默认远程(server)套接字被绑定在回环地址上。但是，显式指定的

             bind_address可以用于绑定套接字到指定的地址上。如果不设置

             bind_address或设置为'*'则表示套接字监听在所有网络接口上。

             只有当远程(server)主机的GatewayPorts选项开启时，指定的

             bind_address才能生效。(见sshd_config)。

             如果port值为0，远程主机(server)监听的端口将被动态分配，并且在

             运行时报告给客户端。

     -S ctl_path

             Specifies the location of a control socket for connection shar‐

             ing, or the string 'none' to disable connection sharing.  Refer

             to the description of ControlPath and ControlMaster in

             ssh_config for details.

     -s      请求在远程主机上调用一个子系统(subsystem)。子系统有助于ssh为

              其他程序(如sftp)提供安全传输。子系统由远程命令指定。

     -T      禁止为ssh分配伪终端。

     -t       强制分配伪终端，重复使用该选项'-tt'将进一步强制。

     -V      显示版本号并退出。

     -v      详细模式，将输出debug消息，可用于调试。'-vvv'可更详细。

     -W host:port

             请求客户端上的标准输入和输出通过安全隧道转发到host:port上，该选

             项隐含了'-N','-T',ExitOnForwardFailure和ClearAllForwardings选项。

     -w local_tun[:remote_tun]

             Requests tunnel device forwarding with the specified tun

             devices between the client (local_tun) and the server

             (remote_tun).

             The devices may be specified by numerical ID or the keyword

             'any', which uses the next available tunnel device.  If

             remote_tun is not specified, it defaults to 'any'.  See also

             the Tunnel and TunnelDevice directives in ssh_config.  If the

             Tunnel directive is unset, it is set to the default tunnel mode,

             which is 'point-to-point'.

     -X      Enables X11 forwarding.  This can also be specified on a per-host

             basis in a configuration file.

             X11 forwarding should be enabled with caution.  Users with the

             ability to bypass file permissions on the remote host (for the

             user's X authorization database) can access the local X11 display

             through the forwarded connection.  An attacker may then be able

             to perform activities such as keystroke monitoring.

             For this reason, X11 forwarding is subjected to X11 SECURITY

             extension restrictions by default.  Please refer to the ssh -Y

             option and the ForwardX11Trusted directive in ssh_config for

             more information.

     -x      Disables X11 forwarding.

     -Y      Enables trusted X11 forwarding.  Trusted X11 forwardings are not

             subjected to the X11 SECURITY extension controls.

     -y      使用syslog发送日志信息。默认情况下日志信息发送到标准错误输出



     除了从命令行获取配置信息，还可以从用户配置文件和全局配置文件中

     获取额外配置信息。详细信息见ssh_config



认证机制

     可用的认证机制及它们的先后顺序为：GSSAPI-based,host-based,public key,

     challenge-response,password。PreferredAuthentications选项可以改变默认的认证顺序



     Host-based authentication works as follows: If the machine the user logs

     in from is listed in /etc/hosts.equiv or /etc/shosts.equiv on the remote

     machine, and the user names are the same on both sides, or if the files

     ~/.rhosts or ~/.shosts exist in the user's home directory on the remote

     machine and contain a line containing the name of the client machine and

     the name of the user on that machine, the user is considered for login.

     Additionally, the server must be able to verify the client's host key

     (see the description of /etc/ssh_known_hosts and ~/.ssh/known_hosts,

     below) for login to be permitted.  This authentication method closes

     security holes due to IP spoofing, DNS spoofing, and routing spoofing.

     [Note to the administrator: /etc/hosts.equiv, ~/.rhosts, and the

     rlogin/rsh protocol in general, are inherently insecure and should be

     disabled if security is desired.]



     公钥认证机制：用户创建公钥/私钥密钥对，将公钥发送给服务端，所以服务端

     知道的是公钥，私钥只有自己知道。



     ~/.ssh/authorized_keys文件列出了允许登录的公钥。当发起连接时，ssh客户端程序

     告诉服务端程序要使用哪个密钥对来完成身份验证，并告诉服务端自己已经访问过

     私钥部分(译者注：不能直接提供私钥给服务端进行比对监测，因为私钥不能泄露)，

     然后服务端则检查对应的公钥部分以确定是否要接受该客户端的连接。



     用户使用ssh-keygen创建密钥对(以rsa算法为例)，将保存在~/.ssh/id_rsa和~/.ssh/id_rsa.pub。

     然后该用户拷贝公钥文件到远程主机上某用户(如A)家目录下的~/.ssh/authorized_keys，

     之后用户就可以以用户A的身份登录到远程主机上。



     公钥认证机制的一种变体是证书认证：只有被信任的证书才允许连接。详细信息见

     ssh-keygen的CERTIFICATES段说明。



     使用公钥认证机制或证书认证机制最方便的方法是'认证代理'，

     详细信息见ssh-agent和ssh_config中的AddKeysToAgent指令段。



     Challenge-response authentication works as follows: The server sends an

     arbitrary 'challenge' text, and prompts for a response.  Examples of

     challenge-response authentication include BSD Authentication (see

     login.conf) and PAM (some non-OpenBSD systems).



     最后，如果所有认证方法都失败，将提示输入密码。输入的密码将被加密传送，

     然后被服务端检测是否正确。



     SSH客户端自动维护和检查一个主机认证信息数据库，所有已知的主机公钥都会

     记录到此文件中。主机信息条目(host key)存放在~/.ssh/known_hosts文件中。

     另外，在检查host key时，/etc/ssh_known_hosts也会被自动检测。

     当host key被改变时，ssh将发出警告，并禁止密钥认证机制以防止服务端欺骗

     或中间人攻击。选项StrictHostKeyChecking选项可用于控制登录时那些未知host key

     如何处理。



     当客户端被服务端接受，服务段将以非交互会话执行给定的命令，若没有给定命令，

     则登录到服务端，并进入到交互会话模式，同时会为登录的用户分配shell，之后

     所有的交互信息都将被加密传输。



     ssh默认会请求交互式会话，这将请求一个伪终端(pty)，使用'-T'或'-t'选项可以

     改变该行为，'-T'是禁止分配伪终端，'-t'则是强制分配伪终端，可使用'-tt'

     表示进一步强制。



     如果为ssh分配了伪终端，则用户可以在此伪终端中使用逃逸字符实现特殊控制。



     如果未分配伪终端给ssh，则连接会话是透明的，可以用来可靠传输二进制数据。

     如果设置逃逸字符为'none'，将使得会话透明，即使它使用了tty终端。

     当命令结束或shell退出时将终止会话连接，所有的X11和TCP连接也都被关闭。



逃逸字符

     当分配了伪终端时，ssh支持一系列的逃逸字符实现特殊功能。

     默认的逃逸首字符为'~'，其后可跟某些特定字符(如下列出)，逃逸字符必须放

     在行尾以实现特定的中断。可在配置文件中使用EscapeChar指令或命令行的'-e'

     选项来改变逃逸首字符。

     ~.      禁止连接

     ~^Z     将ssh放入后台

     ~#      列出已转发的连接

     ~&      Background ssh at logout when waiting for forwarded connection /

               X11 sessions to terminate.

     ~?      列出逃逸字符列表

     ~B      发送BREAK信号给远程主机

     ~C      打开命令行。Open command line.  Currently this allows the addition of port

             forwardings using the -L, -R and -D options (see above).  It also

             allows the cancellation of existing port-forwardings with

             -KL[bind_address:]port for local, -KR[bind_address:]port for

             remote and -KD[bind_address:]port for dynamic port-forwardings.

             !command allows the user to execute a local command if the

             PermitLocalCommand option is enabled in ssh_config.  Basic

             help is available, using the -h option.

     ~R      请求该会话进行密钥更新

     ~V      当错误被写入到stderr时，降低信息的详细程度(loglevel)

     ~v      当错误被写入到stderr时，增加信息的详细程度



TCP转发

     可在配置文件或命令行选项上开启基于安全隧道的任意TCP连接转发功能。

     一个TCP转发可能的应用场景是为了安全连接到邮件服务器，其他场景则主要

     是为了穿过防火墙。



     下面的例子中，建立了IRC客户端和服务端的加密连接，尽管IRC服务端不直

     接支持加密连接。用户在本地指定一个用于转发到远程服务器上的端口，这

     样在本地主机上将开启一个加密的服务，当连接到本地转发端口时，ssh将

     加密和转发此连接。



     下面的示例中，从客户端主机'127.0.0.1'到'server.example.com'的连接将

     使用隧道技术。



         $ ssh -f -L 1234:localhost:6667 server.example.com sleep 10

         $ irc -c '#users' -p 1234 pinky 127.0.0.1



     这个隧道建立在本地和'server.example.com'之间，隧道传递的内容有:

     '#users','pinky',using port 1234. 无论使用的是什么端口，只要大于

     1023(只有root可以在特权端口上建立套接字)，即使端口已被使用也不

     会发生冲突。连接将被转发到远程主机的6667端口上，因为IRC服务的

     默认端口为6667。



     '-f'选项将ssh放入后台，而远程命令'sleep 10'则表示在一段时间(10秒)

     内的连接将通过隧道传输。如果在10秒内没有连接，则ssh退出。

     (也就是说该隧道只在后台保持10秒钟。)



X11 FORWARDING

     If the ForwardX11 variable is set to 'yes' (or see the description of

     the -X, -x, and -Y options above) and the user is using X11 (the DISPLAY

     environment variable is set), the connection to the X11 display is auto‐

     matically forwarded to the remote side in such a way that any X11 pro‐

     grams started from the shell (or command) will go through the encrypted

     channel, and the connection to the real X server will be made from the

     local machine.  The user should not manually set DISPLAY.  Forwarding of

     X11 connections can be configured on the command line or in configuration

     files.



     The DISPLAY value set by ssh will point to the server machine, but with a

     display number greater than zero.  This is normal, and happens because

     ssh creates a 'proxy' X server on the server machine for forwarding the

     connections over the encrypted channel.



     ssh will also automatically set up Xauthority data on the server machine.

     For this purpose, it will generate a random authorization cookie, store

     it in Xauthority on the server, and verify that any forwarded connections

     carry this cookie and replace it by the real cookie when the connection

     is opened.  The real authentication cookie is never sent to the server

     machine (and no cookies are sent in the plain).



     If the ForwardAgent variable is set to 'yes' (or see the description of

     the -A and -a options above) and the user is using an authentication

     agent, the connection to the agent is automatically forwarded to the

     remote side.



VERIFYING HOST KEYS

     当用户第一次连接到一个服务端，将输出服务端公钥的指纹(fingerprint)给用户

     (除非StrictHostKeyChecking配置被禁用了)。这些指纹可通过ssh-keygen来计算。



           $ ssh-keygen -l -f /etc/ssh/ssh_host_rsa_key



     如果某指纹已经存在，可决定对应的密钥是接受还是拒绝。如果仅能获取到服

     务端的传统指纹(MD5)，ssh-keygen的'-E'选项可能会将指纹降级以做指纹匹配。



     由于仅通过查找指纹来比较host key比较困难，所以也支持使用随机数的方式

     可视化比较host key。通过设置VisualHostKey选项为'yes'，客户端连接服务

     端时将显示一小段ASCII图形信息(即图形化的指纹)，无论会话是否是需要交互

     的。通过比较已生成的图形指纹，用户可以轻松地找出host key是否发生了改

     变。但是，由于图形指纹不是很明了，所以相似的图形指纹并不能保证host key

     是没有改变过的，只不过通过图形指纹的方式提供了一个比较好的比较方式。



     要获取所有已知主机(known host)的图形指纹列表，使用下面的命令：



           $ ssh-keygen -lv -f ~/.ssh/known_hosts



     如果指纹是未知的，有一种方法可以验证它：使用DNS。可在DNS的区域文件中添

     加资源记录SSHFP，这样客户端就可以匹配那些已存在的主机指纹。



     在下面的例子中，将使用客户端连接到服务端'host.example.com'。但在此之前，

     应该先将'host.example.com'的SSHFP资源记录添加到DNS区域文件中：



           $ ssh-keygen -r host.example.com.



     将上面命令的输出结果添加到区域文件中。可以检查该资源记录是否可解析：



           $ dig -t SSHFP host.example.com



     最后使用客户端去连接服务端:



           $ ssh -o 'VerifyHostKeyDNS ask' host.example.com

           [...]

           Matching host key fingerprint found in DNS.

           Are you sure you want to continue connecting (yes/no)?



     更多信息请查看ssh_config的VerifyHostKeyDNS选项说明段。



SSH-BASED VIRTUAL PRIVATE NETWORKS

     The following example would connect client network 10.0.50.0/24 with

     remote network 10.0.99.0/24 using a point-to-point connection from

     10.1.1.1 to 10.1.1.2, provided that the SSH server running on the gateway

     to the remote network, at 192.168.1.15, allows it.



     on client:



           # ssh -f -w 0:1 192.168.1.15 true

           # ifconfig tun0 10.1.1.1 10.1.1.2 netmask 255.255.255.252

           # route add 10.0.99.0/24 10.1.1.2



     on server:



           # ifconfig tun1 10.1.1.2 10.1.1.1 netmask 255.255.255.252

           # route add 10.0.50.0/24 10.1.1.1



     Client access may be more finely tuned via the /root/.ssh/authorized_keys

     file (see below) and the PermitRootLogin server option.  The following

     entry would permit connections on tun device 1 from user 'jane' and

     on tun device 2 from user 'john', if PermitRootLogin is set to

     'forced-commands-only':



       tunnel='1',command='sh /etc/netstart tun1' ssh-rsa ... jane

       tunnel='2',command='sh /etc/netstart tun2' ssh-rsa ... john



     Since an SSH-based setup entails a fair amount of overhead, it may be

     more suited to temporary setups, such as for wireless VPNs.  More perma‐

     nent VPNs are better provided by tools such as ipsecctl and

     isakmpd.



ENVIRONMENT

     ssh will normally set the following environment variables:



     DISPLAY               The DISPLAY variable indicates the location of the

                           X11 server.  It is automatically set by ssh to

                           point to a value of the form 'hostname:n', where

                           'hostname' indicates the host where the shell

                           runs, and ‘n’ is an integer ≥ 1.  ssh uses this

                           special value to forward X11 connections over the

                           secure channel.  The user should normally not set

                           DISPLAY explicitly, as that will render the X11

                           connection insecure (and will require the user to

                           manually copy any required authorization cookies).



     HOME                  Set to the path of the user's home directory.



     LOGNAME               Synonym for USER; set for compatibility with sys‐

                           tems that use this variable.



     MAIL                  Set to the path of the user's mailbox.



     PATH                  Set to the default PATH, as specified when compil‐

                           ing ssh.



     SSH_ASKPASS           If ssh needs a passphrase, it will read the

                           passphrase from the current terminal if it was run

                           from a terminal.  If ssh does not have a terminal

                           associated with it but DISPLAY and SSH_ASKPASS are

                           set, it will execute the program specified by

                           SSH_ASKPASS and open an X11 window to read the

                           passphrase.  This is particularly useful when

                           calling ssh from a .xsession or related script.

                           (Note that on some machines it may be necessary to

                           redirect the input from /dev/null to make this

                           work.)



     SSH_AUTH_SOCK         Identifies the path of a UNIX-domain socket used to

                           communicate with the agent.



     SSH_CONNECTION        Identifies the client and server ends of the con‐

                           nection.  The variable contains four space-sepa‐

                           rated values: client IP address, client port num‐

                           ber, server IP address, and server port number.



     SSH_ORIGINAL_COMMAND  This variable contains the original command line if

                           a forced command is executed.  It can be used to

                           extract the original arguments.



     SSH_TTY               This is set to the name of the tty (path to the

                           device) associated with the current shell or com‐

                           mand.  If the current session has no tty, this

                           variable is not set.



     TZ                    This variable is set to indicate the present time

                           zone if it was set when the daemon was started

                           (i.e. the daemon passes the value on to new con‐

                           nections).



     USER                  Set to the name of the user logging in.



     Additionally, ssh reads ~/.ssh/environment, and adds lines of the format

     'VARNAME=value' to the environment if the file exists and users are

     allowed to change their environment.  For more information, see the

     PermitUserEnvironment option in sshd_config.



FILES

     ~/.rhosts

             这个文件用于基于主机的认证机制(见上文)，里面列出允许登录的

             主机/用户对。该文件属主必须是这个对应的用户，且其它用户不

             能有写权限。但如果用户家目录位于NFS分区上时，该文件要求全

             局可读，因为sshd使用root身份读取该文件。大多数情况下，

             推荐权限为'600'。

     ~/.shosts

             该文件的用法与'.rhosts'完全一样，但允许基于主机认证的同时

             禁止使用'rlogin/rsh'登录。

     ~/.ssh/

             该目录是所有用户配置文件和用户认证信息的默认放置目录。虽然

             没有规定要保证该目录中内容的安全，但推荐其内文件只对所有者

             有读/写/执行权限，对其他人完全拒绝。

     ~/.ssh/authorized_keys

             该文件列出了可以用来登录的用户的公钥(DSA,ECDSA,Ed25519,RSA)。

             在sshd的man文档中描述了该文件的格式。该文件不需要高安全性，

             但推荐只有其所有者有读/写权限，对其他人完全拒绝。

     ~/.ssh/config

             该文件是ssh的用户配置文件。在ssh_config的man文档中描述了该

             文件的格式。由于可能会滥用该文件，该文件有严格的权限要求：只

             对所有者有读/写权限，对其他人完全拒绝写权限。

     ~/.ssh/environment

             包含了额外定义的环境变量。见上文ENVIRONMENT。

     ~/.ssh/identity

     ~/.ssh/id_dsa

     ~/.ssh/id_ecdsa

     ~/.ssh/id_ed25519

     ~/.ssh/id_rsa

             包含了认证的私钥。这些文件包含了敏感数据，应该只对所有者可读，

             并拒绝其他人的所有权限(rwx)。如果该文件可被其他人访问，则ssh

             会忽略该文件。可以在生产密钥文件的时候指定passphrase使用3DES

             算法加密该文件。

     ~/.ssh/identity.pub

     ~/.ssh/id_dsa.pub

     ~/.ssh/id_ecdsa.pub

     ~/.ssh/id_ed25519.pub

     ~/.ssh/id_rsa.pub

             包含了认证时的公钥。这些文件中的数据不敏感，允许任何人读取。

     ~/.ssh/known_hosts

             包含了所有已知主机的host key列表。该文件的详细格式见sshd。

     ~/.ssh/rc

             该文件包含了用户使用ssh登录成功，但启用shell(或指定命令执行)

             之前执行的命令。详细信息见sshd的man文档。

             (译者注：也就是说，登录成功后做的第一件事就是执行该文件中的

             命令)

     /etc/ssh/hosts.equiv

             该文件是基于主机认证的文件(见上文)。应该只能让root有写权限。

     /etc/ssh/shosts.equiv

             用法等同于'hosts.equiv'，但允许基于主机认证的同时禁止使用

             'rlogin/rsh'登录。

     /etc/ssh/ssh_config

             ssh的全局配置文件。该文件的格式和选项信息见ssh_config。

     /etc/ssh/ssh_host_key

     /etc/ssh/ssh_host_dsa_key

     /etc/ssh/ssh_host_ecdsa_key

     /etc/ssh/ssh_host_ed25519_key

     /etc/ssh/ssh_host_rsa_key

             这些文件包含了host key的私密部分信息，它们用于基于主机认证。

             (译者注：服务端生成的私钥，主机验证时会将对应公钥存入到客户

               端的known_hosts文件中，这些文件在sshd服务重启时会自动生成)

     /etc/ssh/ssh_known_hosts

             已知host key的全局列表文件。该文件中要包含的host key应该由

             系统管理员准备好。该文件应该要全局可读。详细信息见sshd。

     /etc/ssh/rc

             等同于~/.ssh/rc文件，包含了用户使用ssh登录成功，但启用shell

             (或指定命令执行)之前执行的命令。详细信息见sshd的man文档。

             (译者注：也就是说，登录成功后做的第一件事就是执行该文件中的

             命令)



退出状态码

     ssh将以远程命令执行结果为状态码退出，或者出现错误时以255状态码退出。