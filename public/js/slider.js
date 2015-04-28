$(document).ready( function() {

    // hiji-kakudo-slider
	$( "#hiji-kakudo-slider" ).slider({
      value: 0,
      min: 0,
      max: 150,
      step: 1,
      slide: function( event, ui ) {
        $( "#hiji-kakudo" ).val( ui.value );
      }
    });
    $( "#hiji-kakudo" ).val( $( "#hiji-kakudo-slider" ).slider( "value" ) );

    // ude-furi-slider
    $( "#ude-furi-slider" ).slider({
      value: 0,
      min: 0,
      max: 90,
      step: 1,
      slide: function( event, ui ) {
        $( "#ude-furi" ).val( ui.value );
      }
    });
    $( "#ude-furi" ).val( $( "#ude-furi-slider" ).slider( "value" ) );

    // hiza-mage-slider
	$( "#hiza-mage-slider" ).slider({
      value: 0,
      min: 0,
      max: 150,
      step: 1,
      slide: function( event, ui ) {
        $( "#hiza-mage" ).val( ui.value );
      }
    });
    $( "#hiza-mage" ).val( $( "#hiza-mage-slider" ).slider( "value" ) );

    // asi-furi-slider
    $( "#asi-furi-slider" ).slider({
      value: 0,
      min: 0,
      max: 90,
      step: 1,
      slide: function( event, ui ) {
        $( "#asi-furi" ).val( ui.value );
      }
    });
    $( "#asi-furi" ).val( $( "#asi-furi-slider" ).slider( "value" ) );

    // koshi-kakudo-slider
    $( "#koshi-kakudo-slider" ).slider({
      value: 0,
      min: 0,
      max: 60,
      step: 1,
      slide: function( event, ui ) {
        $( "#koshi-kakudo" ).val( ui.value );
      }
    });
    $( "#koshi-kakudo" ).val( $( "#koshi-kakudo-slider" ).slider( "value" ) );

    /*
    // sokudo-slider
    $( "#sokudo-slider" ).slider({
      value: 4.0,
      min: 1.0,
      max: 7.0,
      step: 0.1,
      slide: function( event, ui ) {
        $( "#sokudo" ).val( ui.value );
      }
    });
    $( "#sokudo" ).val( $( "#sokudo-slider" ).slider( "value" ) );
    */

} );
