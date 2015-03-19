### protocol

####push

    push a message into a channel,if channel not existed,create one.

    format:  push [channel] [size]\r\n[body]
    example: push asdf 2\r\naa
    return:
             [channel] ok
             [channel] err [err msg]

####pull
    
    pull a msg from channel,if channel not existed,create one.
    if timeout is not specified, pull will wait forever.

    format: pull [channel] [(optional)timeout]\r\n
    example: pull asdf\r\n
             pull asdf 5\r\n
    return:
            [channel] msg [id] [timestramp] [size]\r\n[body]
            [channel] err [err msg]


####commit

    commit a msg on a channel,so we can safely remove a msg from channel.

    format: commit [channel]\r\n
    example: commit asdf

    return:
            [channel] ok
            [channel] err [err msg]
    




