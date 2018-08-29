module Main.ViewMobileNav exposing (bottomLinksAStyle, bottomLinksStyle, closeStyle, logoStyle, mobileNavLIStyle, navAStyle, navColor, navIconStyle, navStyle, render, ulStyle)

import Html exposing (..)
import Html.Attributes exposing (attribute, href, style)
import Html.Events exposing (onClick)
import Main.Model exposing (Model)
import Main.Msg exposing (Msg(..))
import Main.Routing
    exposing
        ( Route(..)
        )
import Material.Icon as Icon
import Material.List as Lists
import Material.Options as Options


render : Model -> Html Msg
render model =
    if not model.showMobileNav then
        div [] []

    else
        div
            [ navStyle ]
            [ div [ closeStyle, onClick ToggleMobileNav ] [ Icon.view "close" [ Icon.size36 ] ]
            , div [ logoStyle ] [ text "Menu" ]
            , Lists.ul [ ulStyle ]
                [ Lists.li [ mobileNavLIStyle ]
                    [ Lists.content []
                        [ Lists.icon "account_box" [ navIconStyle ]
                        , a
                            [ navAStyle
                            , href ""
                            , navAStyle
                            ]
                            [ text "Account" ]
                        ]
                    ]
                , Lists.li [ mobileNavLIStyle ]
                    [ Lists.content []
                        [ div
                            [ navAStyle
                            ]
                            [ Lists.icon "vpn_key" [ navIconStyle ]
                            , text "Logout"
                            ]
                        ]
                    ]
                ]
            , div [ bottomLinksStyle ]
                []
            ]


navStyle : Attribute a
navStyle =
    style
        [ ( "position", "fixed" )
        , ( "top", "0" )
        , ( "bottom", "0" )
        , ( "left", "0" )
        , ( "width", "100%" )
        , ( "height", "100%" )
        , ( "background", "#fff" )
        , ( "z-index", "100" )
        ]


ulStyle =
    Options.attribute <| style [ ( "margin-top", "60px" ) ]


closeStyle : Attribute a
closeStyle =
    style
        [ ( "position", "absolute" )
        , ( "top", "20px" )
        , ( "right", "20px" )
        , ( "cursor", "pointer" )
        , ( "color", navColor )
        ]


logoStyle =
    style [ ( "padding-top", "30px" ), ( "padding-left", "20px" ) ]


mobileNavLIStyle =
    Options.attribute <| style [ ( "border-bottom", "1px solid #eee" ) ]


navIconStyle =
    Options.attribute <| style [ ( "color", navColor ) ]


navAStyle =
    style
        [ ( "color", "#000" )
        , ( "text-decoration", "none" )
        , ( "font-weight", "normal" )
        , ( "line-height", "18px" )
        , ( "font-size", "12px" )
        , ( "letter-spacing", "1px" )
        , ( "color", navColor )
        ]


bottomLinksStyle =
    style
        [ ( "position", "absolute" )
        , ( "bottom", "0" )
        , ( "left", "0" )
        , ( "padding", "10px" )
        , ( "text-align", "center" )
        , ( "width", "100%" )
        ]


bottomLinksAStyle =
    style
        [ ( "text-decoration", "none" )
        , ( "color", navColor )
        , ( "font-size", "12px" )
        , ( "padding", "10px 15px" )
        , ( "display", "inline-block" )
        ]


navColor =
    "#9d9cac"
