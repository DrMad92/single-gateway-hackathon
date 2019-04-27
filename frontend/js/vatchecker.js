$(document).ready(function(){
    $("#country-div-id, #product-div-id").hide();
    var countryData = {};
    var categoriesList = [];
    $("select#country-select-id").change(function(){
        $("#country-div-id").fadeIn();
        $("#product-div-id").fadeOut();
        var selectedCountry = $(this).children("option:selected").val();
        var data = {
            country: selectedCountry
        }
        $.ajax({
            method: "POST",
            url: "http://176.112.147.202:9999",
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
                $("#default-vat-result-id").html("<b>Standard VAT</b>: " + countryData.standardRate + " %");
                $("#treshold-vat-result-id").html("<b>Threshold</b>: " + countryData.threshold + " \u20AC");
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
        $("#product-comment-result-id").html('');
        $("#product-description-result-id").html('');
        $("#product-vat-result-id").html('');
        $("#product-div-id").fadeIn();
        var selectedProduct = $(this).children("option:selected").val();
        if (selectedProduct == "other") {
            $("#product-vat-result-id").html('<span><b>VAT</b>: ' + countryData.standardRate + ' %</span><br>');
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
                $("#product-comment-result-id").html('<b>Comment</b>: ' + category.comments);
            }
            $("#product-description-result-id").html('<b>Description</b>: ' + category.description);
            $("#product-vat-result-id").html('<b>' + beautify(category.name) + ' VAT</b>: ' + category.reducedRate + ' %');
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