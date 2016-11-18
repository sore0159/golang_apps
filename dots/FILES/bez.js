var canvas = document.getElementById("can");
var pic = document.getElementById("pic");
var pointList = [];
var cMouseDown = false;
var selectedButton = -1;

canvas.oncontextmenu= function() { return false; };
canvas.addEventListener("mousemove", dragButton);
canvas.addEventListener("mousedown", function() {
    cMouseDown = true;
});
canvas.addEventListener("mouseup", function() {
    cMouseDown = false;
    selectedButton = -1;
});
canvas.addEventListener("mouseleave", function() {
    cMouseDown = false;
    selectedButton = -1;
});

function dragButton(event) {
    if (!cMouseDown) {
        return;
    }
    var clickx = event.pageX - canvas.offsetLeft;
    var clicky = event.pageY - canvas.offsetTop;
    var flag;
    if (selectedButton > -1) {
        pointList[selectedButton][0] = clickx;
        pointList[selectedButton][1] = clicky;
        drawBez(pointList);
        return;
    }
    pointList.forEach(function(pt, i) {
        if (flag) {
            return;
        }
        var dx = pt[0] - clickx;
        var dy = pt[1] - clicky;
        if (dx*dx + dy*dy < 81) {
            pt[0] = clickx;
            pt[1] = clicky;
            drawBez(pointList);
            selectedButton = i;
            flag = true;
        }
    });
}

for (var i = 0; i < 4; i += 1) {
        pt = [Math.random()*500, Math.random()*500];
        pointList.push(pt);
}
putBez(pointList);
drawBez(pointList);

function putBez(pts) {
    fetch("bez.api", { method: 'POST', credentials: 'include', body: JSON.stringify(pointList) }).then(function(response) {
    pic.src = "bez.png?"+ new Date().getTime();
});
}


function drawBez(pts) {
    var ctx = canvas.getContext("2d");
    ctx.clearRect(0,0,500,500);
    var gc = new Path2D();
    if (pts.length === 3) {
		gc.moveTo(pts[0][0], pts[0][1]);
		gc.lineTo(pts[1][0], pts[1][1]);
		gc.lineTo(pts[2][0], pts[2][1]);
		ctx.lineWidth = 2;
		ctx.strokeStyle = "#AAA";
		ctx.stroke(gc);
		//
        gc = new Path2D();
		gc.moveTo(pts[0][0], pts[0][1]);
		gc.quadraticCurveTo(pts[1][0], pts[1][1], pts[2][0], pts[2][1]);
		ctx.strokeStyle = "#000";
		ctx.lineWidth = 5;
		ctx.stroke(gc);
		//
        ctx.fillStyle = "#F00";
		ctx.lineWidth = 1;
		blip(pts[0], ctx);
		blip(pts[2], ctx);
        ctx.fillStyle = "#00F";
		blip(pts[1], ctx);
    } else if (pts.length === 4) {
		gc.moveTo(pts[0][0], pts[0][1]);
		gc.lineTo(pts[1][0], pts[1][1]);
		gc.moveTo(pts[3][0], pts[3][1]);
		gc.lineTo(pts[2][0], pts[2][1]);
		ctx.lineWidth = 2;
		ctx.strokeStyle = "#AAA";
		ctx.stroke(gc);
		//
        gc = new Path2D();
		ctx.strokeStyle = "#000";
		ctx.lineWidth = 5;
		gc.moveTo(pts[0][0], pts[0][1]);
		gc.bezierCurveTo(pts[1][0], pts[1][1], pts[2][0], pts[2][1], pts[3][0], pts[3][1]);
        ctx.stroke(gc);
		//
		ctx.lineWidth = 1;
        ctx.fillStyle = "#F00";
		blip(pts[0], ctx);
		blip(pts[3], ctx);
        ctx.fillStyle = "#00F";
		blip(pts[1], ctx);
		blip(pts[2], ctx);
    }
}
function blip(pt, ctx) {
    var path = new Path2D();
	var r = 8;
	path.arc(pt[0], pt[1], r, 0, 2*Math.PI);
	ctx.fill(path);
    ctx.stroke(path);
}

var subButton = document.getElementById("subButton");
subButton.addEventListener("mouseup", function() {
    putBez(pointList);
});
var changeButton = document.getElementById("changeButton");
changeButton.addEventListener("mouseup", function() {
    var pt;
    if (pointList.length === 4) {
        pt = pointList.pop();
        pointList.pop();
        pointList.push(pt);
        drawBez(pointList);
        changeButton.innerText = "Cubic";
    } else if (pointList.length < 4) {
        pt = pointList.pop();
        pointList.push([Math.random()*500, Math.random()*500]);
        pointList.push(pt);
        drawBez(pointList);
        changeButton.innerText = "Quadratic";
    }
});
