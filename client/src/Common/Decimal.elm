module Common.Decimal exposing
    ( getDecimalZeros
    , getDecimals
    , renderDecimal
    , renderDecimalWithPrecision
    , round
    , round8
    , splitPrice
    )

import Html exposing (..)
import Html.Attributes exposing (style)


round : Int -> String -> String
round precision value =
    case String.split "." value of
        [ number, decimals ] ->
            number
                ++ "."
                ++ (if String.length decimals > precision then
                        String.left precision decimals

                    else
                        decimals
                   )

        _ ->
            value


{-| rounds a decimal to the 8th precision
-}
round8 : String -> String
round8 value =
    round 8 value


renderDecimalWithPrecision : String -> Int -> Html a
renderDecimalWithPrecision value precision =
    case String.contains "." value of
        False ->
            span [] [ text value ]

        True ->
            let
                ( p, decimals, decimalZeros ) =
                    splitPrice (round precision value) "" ""
            in
            span []
                [ text p
                , span [ style [ ( "font-size", "80%" ) ] ]
                    [ text decimals
                    , span [ style [ ( "opacity", "0.5" ) ] ] [ text decimalZeros ]
                    ]
                ]


renderDecimal : String -> Html a
renderDecimal value =
    renderDecimalWithPrecision value 8


{-| Splits decimal in three parts as follows:
10.340420000 becomes
part1: 10. --- main number
part2: 34042 --- decimal value
part3: 00000 --- Tail decimal zeros
-}
splitPrice : String -> String -> String -> ( String, String, String )
splitPrice value decimals decimalZeros =
    let
        ( decimalZeros, remaining_ ) =
            getDecimalZeros value ""

        ( decimals, remaining ) =
            getDecimals remaining_ ""
    in
    ( remaining, decimals, decimalZeros )


getDecimalZeros : String -> String -> ( String, String )
getDecimalZeros value decimalZeros =
    case String.right 1 value of
        "0" ->
            getDecimalZeros (String.dropRight 1 value) (decimalZeros ++ "0")

        remaining ->
            ( decimalZeros, value )


getDecimals : String -> String -> ( String, String )
getDecimals value decimals =
    case String.right 1 value of
        "." ->
            ( decimals, value )

        digit ->
            getDecimals (String.dropRight 1 value) (digit ++ decimals)
