module Actions.Sub exposing (subscriptions)

import Actions.Msg exposing (Msg(..))
import Main.Context exposing (Context)
import Time


subscriptions : Context -> Sub Msg
subscriptions ctx =
    Sub.batch
        [ Time.every 500 (always TickerTimout)
        ]



-- Sub.none
-- Time.every 500 (always TickerTimout)
