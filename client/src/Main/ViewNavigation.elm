module Main.ViewNavigation exposing (render)

import Html exposing (..)
import Html.Attributes exposing (attribute, href, style)
import Html.Events exposing (onClick)
import Main.Model exposing (Model)
import Main.Msg exposing (Msg(..))
import Main.Routing
    exposing
        ( Route(..)
        , actionsPath
        , homepagePath
        , newActionPath
        , portfolioPath
        , tokensPath
        )
import Material.Icon as Icon


render : Model -> Html Msg
render model =
    div [ bhStyle ]
        [ div []
            [ a
                [ href actionsPath
                , bhNavItemStyle
                , activeRouteStyle (model.context.route == ActionsRoute)
                , style [ ( "width", "33%" ) ]
                ]
                [ Icon.view "home"
                    [ Icon.size36
                    ]
                ]
            , a
                [ href newActionPath
                , bhNavItemStyle
                , activeRouteStyle (model.context.route == CreateActionRoute)
                , style [ ( "width", "33%" ) ]
                ]
                [ Icon.view "add_circle_outline"
                    [ Icon.size36
                    ]
                ]
            , a
                [ href portfolioPath
                , bhNavItemStyle
                , activeRouteStyle (model.context.route == PortfolioRoute)
                , style [ ( "width", "34%" ) ]
                ]
                [ Icon.view "account_box"
                    [ Icon.size36
                    ]
                ]
            ]
        ]



-- Common


activeRouteStyle isActive =
    case isActive of
        True ->
            style
                [ ( "color", "#ff5722" )
                ]

        False ->
            style
                [ ( "color", "inherit" )
                ]


brandColor : String
brandColor =
    "#3f51b5"


textColor : String
textColor =
    "#e1ecf2"


thStyle : Attribute a
thStyle =
    style
        [ ( "height", "45px" )
        , ( "padding-top", "10px" )
        , ( "box-sizing", "border-box" )
        , ( "padding-left", "10px" )
        , ( "padding-right", "10px" )
        , ( "line-height", "25px" )
        , ( "background", brandColor )
        , ( "width", "100%" )
        , ( "position", "fixed" )
        , ( "top", "0" )
        , ( "z-index", "1000" )
        ]


thNavStyle : Attribute a
thNavStyle =
    style
        [ ( "float", "right" )
        ]


thNavItemStyle : Attribute a
thNavItemStyle =
    style
        [ ( "margin", "0 8px" )
        , ( "text-decoration", "none" )
        , ( "font-size", "14px" )
        , ( "letter-spacing", ".7px" )
        , ( "color", textColor )
        ]


thSubNavItemStyle : Attribute a
thSubNavItemStyle =
    style
        [ ( "padding", "10px" )
        , ( "text-decoration", "none" )
        , ( "font-size", "12px" )
        , ( "font-weight", "normal" )
        , ( "color", textColor )
        ]


thLogoStyle : Attribute a
thLogoStyle =
    style
        [ ( "text-transform", "uppercase" )
        , ( "text-decoration", "none" )
        , ( "font-size", "12px" )
        , ( "margin", "0px 40px 0px 25px" )
        , ( "color", textColor )
        ]



-- Mobile Styles, BottomHeader :bh


bhStyle : Attribute a
bhStyle =
    style
        [ ( "position", "fixed" )
        , ( "bottom", "0" )
        , ( "left", "0" )
        , ( "width", "100%" )
        , ( "height", "60px" )
        , ( "background", "white" )
        , ( "min-width", "320px" )
        , ( "opacity", ".95" )
        , ( "z-index", "100" )
        , ( "border-top", "1px solid #ddd" )
        ]


bhNavItemStyle : Attribute a
bhNavItemStyle =
    style
        [ ( "display", "inline-block" )
        , ( "text-align", "center" )
        , ( "line-height", "50px" )
        , ( "text-decoration", "none" )
        , ( "color", "black" )
        , ( "overflow", "hidden" )
        , ( "cursor", "pointer" )
        , ( "padding-top", "8px" )
        ]
