module Common.Json exposing (..)

import Json.Decode as JD
import Json.Decode.Pipeline as JP
import Debug


type alias EmptyResponse =
    {}


emptyResponseDecoder =
    JD.succeed {}


decodeAt : String -> String -> JD.Decoder a -> Result String a
decodeAt paramName json decoder =
    JD.decodeString
        (JD.at [ "p" ] <| JD.at [ paramName ] <| decoder)
        json


deocdeIntWithDefault : Int -> String -> String -> Int
deocdeIntWithDefault default paramName json =
    case decodeAt paramName json JD.int of
        Ok intv ->
            intv

        Err err ->
            Debug.log err
                default
