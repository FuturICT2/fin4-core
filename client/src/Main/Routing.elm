module Main.Routing exposing
    ( Route(..)
    , fastSignupPath
    , forgotPassPath
    , homepagePath
    , loginPath
    , matchers
    , newTokenPath
    , parseLocation
    , portfolioPath
    , signupPath
    , termsPath
    , tokensPath
    )

import Navigation exposing (Location)
import UrlParser exposing ((</>), Parser, int, map, oneOf, parseHash, s, string, top)


type Route
    = NotFoundRoute
    | HomepageRoute
    | TokenRoute Int
    | PortfolioRoute
    | TokensRoute
    | CreateTokenRoute
    | UserLoginRoute
    | UserForgotRoute
    | UserForgotResetPassRoute Int String
    | UserSignupRoute
    | UserFastSignupRoute


matchers : Parser (Route -> a) a
matchers =
    oneOf
        [ map HomepageRoute top
        , map TokenRoute (s "token" </> int)
        , map PortfolioRoute (s "portfolio")
        , map TokensRoute (s "tokens")
        , map CreateTokenRoute (s "new")
        , map UserLoginRoute (s "login")
        , map UserSignupRoute (s "signup")
        , map UserFastSignupRoute (s "fsignup")
        , map UserForgotRoute (s "forgot")
        , map UserForgotResetPassRoute (s "forgot-reset-password" </> int </> string)
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


loginPath : String
loginPath =
    "#login"


signupPath : String
signupPath =
    "#signup"


forgotPassPath : String
forgotPassPath =
    "#forgot"


fastSignupPath : String
fastSignupPath =
    "#fsignup"


termsPath : String
termsPath =
    "#site-terms"


privacyPath : String
privacyPath =
    "#site-privacy"
