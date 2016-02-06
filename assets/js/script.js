$(function() {
  function createModal(text) {
    function closeModal(e) {
      e.stopPropagation();

      $('.modal-overlay').remove();
    }

    var container = $('<div>', {
      'class': 'modal-overlay'
    }).append(
      $('<div>', {
        'class': 'modal'
      }).append(
        $('<p>', {
          'class': 'modal-text',
          text: text
        }),
        $('<button>', {
          'class': 'modal-button',
          text: 'CLOSE'
        }).on('click', closeModal)
      )
    ).on('click', closeModal);

    $('body').append(container);
  }

  $('.item-pay').on('click', function(e) {
    e.stopPropagation();

    createModal(
      "Tnx 4 your money, loser!"
    );
  });

  $('.item-click').on('click', function(e) {
    e.stopPropagation();

    createModal(
      "You clicked your lucky button! Give us your money now!"
    );
  });
})
