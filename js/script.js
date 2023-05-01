// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M.
// Created on: Sep 2020
// This file contains the JS functions for index.html

function myButtonClicked() {

  const age = parseInt(document.getElementById('age-entered').value)

  if (age >= 18) {

    document.getElementById('answer').innerHTML = 'You can see an R rated movie alone.'

  } else if (age >= 13) {

    document.getElementById('answer').innerHTML = 'You can see a PG-13 movie alone.'

  } else {
      
      document.getElementById('answer').innerHTML = 'You can see a G rated movie alone.'
  }
}