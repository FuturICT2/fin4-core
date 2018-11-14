'use strict'

import './res/style.scss'
require('./index.html');
const Elm = require('./Main.elm')
const apiBase = '/wapi'

const app = Elm.Main.embed(document.getElementById('app'), {
  apiBase: apiBase
});
