$(document).ready(function(){
    var countryData = {};
    $("select#country-select-id").change(function(){
    var selectedCountry = $(this).children("option:selected").val();
    var data = {
        country: selectedCountry
    }
    $.ajax({
        method: "POST",
        url: "http://localhost:9999",
        dataType: 'json',
        contentType: 'application/json',
        data: JSON.stringify(data),
        error: function(err){
             alert(err);
        },
        success: function(resp)
        {
            countryData = resp;
            $("#product-select-id").empty();
            $("#product-select-id").html('<option id="empty-select" disabled selected value> -- select a product -- </option>\
                                        <option id="other" value="other"> Other </option>');
            $.each(Object.keys(countryData.categories), function(key,value) {
                $("#empty-select").after('<option value="' + value + '">' + beautify(value) + '</option>');
              })
        }
      });
   });
   $("select#product-select-id").change(function(){
        $("#vat-result-id").empty();
        var selectedProduct = $(this).children("option:selected").val();
        if (selectedProduct == "other") {
            $("#vat-result-id").html('<span>VAT: ' + countryData.standardRate + '%</span><br>');
        }
        else {
            $("#vat-result-id").html('<span>VAT: ' + countryData.categories[selectedProduct] + '%</span><br>');
        }
   });
   $(document).ajaxStart(function() {
    $(document.body).css({'cursor' : 'wait'});
    }).ajaxStop(function() {
        $(document.body).css({'cursor' : 'default'});
    });
});

function beautify(string){
    return string.charAt(0).toUpperCase() + string.slice(1).replace(/_/g, ' ');;
}