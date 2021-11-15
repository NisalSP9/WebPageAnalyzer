$(document).ready(function () {
    $("#btn").click(function (e) {
        var validate = Validate();
        var requURL = {
            url: $("#requ_url").val()
        }

        if (validate.length === 0) {
            on();
            $.ajax({
                type: "POST",
                url: "http://localhost:3000/api/gethtml",
                dataType: "json",
                data: JSON.stringify(requURL),
                success: function (result) {
                    var table = $("<table><th>Analyze Report</th>");

                    table.append("<tr><th>HTML Version :</th><td>" + result["HTMLVersion"] + "</td></tr>");
                    table.append("<tr><th>Page Title :</th><td>" + result["pageTitle"] + "</td></tr>");
                    table.append("<tr><th>Headings :</th><td>" + JSON.stringify(result["headings"]) + "</td></tr>");
                    table.append("<tr><th>Internal Links :</th><td>" + result["internal"] + "</td></tr>");
                    table.append("<tr><th>External Links :</th><td>" + result["external"] + "</td></tr>");
                    table.append("<tr><th>Inaccessible Links :</th><td>" + result["inaccessible"] + "</td></tr>");
                    table.append("<tr><th>Has a login form :</th><td>" + result["loginForm"] + "</td></tr>");
                    $("#report").html(table);
                    off()
                },
                error: function (rst) {
                    var table = $("<table><th>Analyze Report</th>");

                    table.append("<tr><th>HTML Version :</th><td></td></tr>");
                    table.append("<tr><th>Page Title :</th><td></td></tr>");
                    table.append("<tr><th>Headings :</th><td></td></tr>");
                    table.append("<tr><th>Internal Links :</th><td></td></tr>");
                    table.append("<tr><th>External Links :</th><td></td></tr>");
                    table.append("<tr><th>Inaccessible Links :</th><td></td></tr>");
                    table.append("<tr><th>Has a login form :</th><td></td></tr>");
                    $("#report").html(table);
                    off()
                    alert(rst["responseText"])
                }
            });
        }
    });

    $(document).ajaxStart(function () {
        $("img").show();
    });

    $(document).ajaxStop(function () {
        $("img").hide();
    });

    function Validate() {
        let errorCode = "";
        if ($("#requ_url").val() === "") {
            errorCode = "ErrorFound";
            alert("Enter URL");
        }
        return errorCode;
    }

    function on() {
        document.getElementById("overlay").style.display = "block";
    }

    function off() {
        document.getElementById("overlay").style.display = "none";
    }
});