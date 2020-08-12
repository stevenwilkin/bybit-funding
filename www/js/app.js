$(function() {
  $.get('/funding', function(data) {
    $('#current').text((data.current * 100) + '%');
    $('#predicted').text((data.predicted * 100) + '%');
  });
});
