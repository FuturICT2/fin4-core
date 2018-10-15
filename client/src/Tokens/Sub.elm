module Tokens.Sub exposing (subscriptions)

import Main.Context exposing (Context)
import Time
import Tokens.Msg exposing (Msg(..))


subscriptions : Context -> Sub Msg
subscriptions ctx =
    Time.every 500 (always TickerTimout)
