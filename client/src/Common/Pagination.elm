module Common.Pagination exposing (..)

import Html exposing (..)
import Html.Attributes exposing (attribute, style)
import Html.Events exposing (onClick)
import Material.Options as Options exposing (css)
import Main.Context exposing (Context)
import Common.Styles exposing (textCenter, textRight)


render ctx onClickMsg allCount page limit =
    let
        pages =
            ceiling <| toFloat allCount / toFloat limit

        bills =
            if ctx.window.isMobile then
                1
            else
                4

        start =
            if (page - bills) < 1 then
                1
            else
                page - bills

        end =
            if page + bills > pages then
                pages + 1
            else
                page + bills + 1
    in
        case pages < 2 of
            True ->
                div [] []

            False ->
                div
                    [ style
                        [ ( "display", "flex" )
                        , ( "flex-direction", "row" )
                        , ( "flex-wrap", "nowrap" )
                        , ( "justify-content", "space-between" )
                        , ( "padding", "10px 0" )
                        ]
                    ]
                    [ div []
                        [ if page == 1 then
                            a [ attribute "class" "page" ] [ text "<<" ]
                          else
                            a [ attribute "class" "page", onClick <| onClickMsg 1 ] [ text "<<" ]
                        , if page == 1 then
                            a [ attribute "class" "page" ] [ text "<" ]
                          else
                            a [ attribute "class" "page", onClick <| onClickMsg (page - 1) ] [ text "<" ]
                        ]
                    , div [ textCenter ] <|
                        List.concat
                            [ pagesTo onClickMsg start page []
                            , [ a [ attribute "class" "page active" ] [ text <| toString page ] ]
                            , pagesTo onClickMsg (page + 1) end []
                            ]
                    , div [ textRight ]
                        [ if page == pages then
                            a [ attribute "class" "page" ] [ text ">" ]
                          else
                            a [ attribute "class" "page", onClick <| onClickMsg (page + 1) ] [ text ">" ]
                        , if page == pages then
                            a [ attribute "class" "page last" ] [ text ">>" ]
                          else
                            a [ attribute "class" "page last", onClick <| onClickMsg (pages) ] [ text ">>" ]
                        ]
                    ]


pagesTo onClickMsg from to result =
    let
        nextFrom =
            from + 1
    in
        if from == to || from > to then
            List.reverse result
        else
            let
                current =
                    a [ attribute "class" "page", onClick <| onClickMsg from ] [ text <| toString from ]
            in
                pagesTo onClickMsg nextFrom to (current :: result)
