$(document).ready(function(){
    var countryData = {};
    var categoriesList = [];
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
            categoriesList = resp.categories;
            $("#default-vat-result-id").html("Standard VAT: " + countryData.standardRate + " %");
            $("#treshold-vat-result-id").html("Threshold: " + countryData.threshold + " \u20AC");
            $("#product-select-id").empty();
            $("#product-select-id").html('<option id="empty-select" disabled selected value> -- select a product -- </option>\
                                        <option id="other" value="other"> Other </option>');
            $.each(categoriesList, function(key,value) {
                $("#empty-select").after('<option value="' + value.name + '">' + beautify(value.name) + '</option>');
              })
        }
      });
   });
   $("select#product-select-id").change(function(){
        $("#vat-result-id").empty();
        var selectedProduct = $(this).children("option:selected").val();
        if (selectedProduct == "other") {
            $("#vat-result-id").html('<span>VAT: ' + countryData.standardRate + ' %</span><br>');
        }
        else {
            var category;
            categoriesList.some(function(item) {
                if (item.name == selectedProduct){
                    category = item;
                    return;
                }
              });
            if (category.comments != ""){
                $("#product-comment-result-id").html('Comment: ' + category.comments);
            }
            $("#product-description-result-id").html('Description: ' + category.description);
            $("#product-vat-result-id").html(beautify(category.name) + ' VAT: ' + category.reducedRate + ' %');
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