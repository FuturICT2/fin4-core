module Asset.Sub exposing (subscriptions)

import Asset.Msg exposing (Msg(..))
import Asset.Ports exposing (assetBlockImageRead)
import Main.Context exposing (Context)
import Time


subscriptions : Context -> Sub Msg
subscriptions ctx =
    Sub.batch
        [ assetBlockImageRead ImageRead
        , Time.every 500 (always TickerTimout)
        ]
