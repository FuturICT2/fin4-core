module Main.Routing exposing
    ( Route(..)
    , homepagePath
    , matchers
    , newTokenPath
    , parseLocation
    , portfolioPath
    , tokensPath
    )

import Navigation exposing (Location)
import UrlParser exposing ((</>), Parser, map, oneOf, parseHash, s, string, top)


type Route
    = NotFoundRoute
    | HomepageRoute
    | TokensRoute
    | PortfolioRoute
    | TokenCreatorRoute


matchers : Parser (Route -> a) a
matchers =
    oneOf
        [ map HomepageRoute top
        , map TokensRoute (s "tokens")
        , map PortfolioRoute (s "portfolio")
        , map TokenCreatorRoute (s "new")
        ]


parseLocation : Location -> Route
parseLocation location =
    case parseHash matchers location of
        Just route ->
            route

        Nothing ->
            NotFoundRoute


homepagePath : String
homepagePath =
    "#"


tokensPath : String
tokensPath =
    "#tokens"


trendingPath : String
trendingPath =
    "#tokens"


newTokenPath : String
newTokenPath =
    "#new"


portfolioPath : String
portfolioPath =
    "#portfolio"
