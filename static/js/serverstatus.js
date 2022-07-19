// serverstatus.js
var error = 0;
var d = 0;
var server_status = new Array();

function getDuration(second) {
	var duration
	var days = Math.floor(second / 86400);
	var hours = Math.floor((second % 86400) / 3600);
	var minutes = Math.floor(((second % 86400) % 3600) / 60);
	var seconds = Math.floor(((second % 86400) % 3600) % 60);
	if(days>0)  duration = days + "天" + hours + "小时" + minutes + "分" + seconds + "秒";
	else if(hours>0)  duration = hours + "小时" + minutes + "分" + seconds + "秒";
	else if(minutes>0) duration = minutes + "分" + seconds + "秒";
	else if(seconds>0) duration = seconds + "秒";
	return duration;
}

function timeSince(date) {
	if(date == 0)
		return "从未.";

	var seconds = Math.floor((new Date() - date) / 1000);
	var interval = Math.floor(seconds / 31536000);

	if (interval > 1)
		return interval + " 年前.";
	interval = Math.floor(seconds / 2592000);
	if (interval > 1)
		return interval + " 月前.";
	interval = Math.floor(seconds / 86400);
	if (interval > 1)
		return interval + " 日前.";
	interval = Math.floor(seconds / 3600);
	if (interval > 1)
		return interval + " 小时前.";
	interval = Math.floor(seconds / 60);
	if (interval > 1)
		return interval + " 分钟前.";
	/*if(Math.floor(seconds) >= 5)
		return Math.floor(seconds) + " seconds";*/
	else
		return "几秒前.";
}

function bytesToSize(bytes, precision, si)
{
	var ret;
	si = typeof si !== 'undefined' ? si : 0;
	if(si != 0) {
		var kilobyte = 1000;
		var megabyte = kilobyte * 1000;
		var gigabyte = megabyte * 1000;
		var terabyte = gigabyte * 1000;
	} else {
		var kilobyte = 1024;
		var megabyte = kilobyte * 1024;
		var gigabyte = megabyte * 1024;
		var terabyte = gigabyte * 1024;
	}

	if ((bytes >= 0) && (bytes < kilobyte)) {
		return bytes + ' B';

	} else if ((bytes >= kilobyte) && (bytes < megabyte)) {
		ret = (bytes / kilobyte).toFixed(precision) + ' K';

	} else if ((bytes >= megabyte) && (bytes < gigabyte)) {
		ret = (bytes / megabyte).toFixed(precision) + ' M';

	} else if ((bytes >= gigabyte) && (bytes < terabyte)) {
		ret = (bytes / gigabyte).toFixed(precision) + ' G';

	} else if (bytes >= terabyte) {
		ret = (bytes / terabyte).toFixed(precision) + ' T';

	} else {
		return bytes + ' B';
	}
	if(si != 0) {
		return ret + 'B';
	} else {
		return ret + 'iB';
	}
}

function uptime() {
	$.getJSON("http://localhost:8085/api/node/allStatus", function(result) {
		$("#loading-notice").remove();
		if(result.reload)
			setTimeout(function() { location.reload(true) }, 1000);

		for (var i = 0, rlen=result.servers.length; i < rlen; i++) {
			var TableRow = $("#servers tr#r" + i);
			var ExpandRow = $("#servers #rt" + i);
			var hack; // fuck CSS for making me do this
			if(i%2) hack="odd"; else hack="even";
			if (!TableRow.length) {
				$("#servers").append(
					"<tr id=\"r" + i + "\" data-toggle=\"collapse\" data-target=\"#rt" + i + "\" class=\"accordion-toggle " + hack + "\">" +
						"<td id=\"online4\"><div class=\"progress\"><div style=\"width: 100%;\" class=\"progress-bar progress-bar-warning\"><small>加载中</small></div></div></td>" +
						"<td id=\"host\">加载中</td>" +
						"<td id=\"ip_status\"><div class=\"progress\"><div style=\"width: 100%;\" class=\"progress-bar progress-bar-warning\"><small>加载中</small></div></div></td>" +
						"<td id=\"name\">加载中</td>" +
						"<td id=\"type\">加载中</td>" +
						"<td id=\"location\">加载中</td>" +
						"<td id=\"uptime\">加载中</td>" +
						"<td id=\"load\">加载中</td>" +
						"<td id=\"network\">加载中</td>" +
						"<td id=\"traffic\">加载中</td>" +
						"<td id=\"cpu\"><div class=\"progress\"><div style=\"width: 100%;\" class=\"progress-bar progress-bar-warning\"><small>加载中</small></div></div></td>" +
						"<td id=\"memory\"><div class=\"progress\"><div style=\"width: 100%;\" class=\"progress-bar progress-bar-warning\"><small>加载中</small></div></div></td>" +
						"<td id=\"hdd\"><div class=\"progress\"><div style=\"width: 100%;\" class=\"progress-bar progress-bar-warning\"><small>加载中</small></div></div></td>" +
						"<td id=\"ping\"><div class=\"progress\"><div style=\"width: 100%;\" class=\"progress-bar progress-bar-warning\"><small>加载中</small></div></div></td>" +
					"</tr>" +
					"<tr class=\"expandRow " + hack + "\"><td colspan=\"16\"><div class=\"accordian-body collapse\" id=\"rt" + i + "\">" +
						"<div id=\"expand_mem\">加载中</div>" +
						"<div id=\"expand_swap\">加载中</div>" +
						"<div id=\"expand_hdd\">加载中</div>" +
						"<div id=\"expand_tupd\">加载中</div>" +
						"<div id=\"expand_ping\">加载中</div>" +
						"<div id=\"expand_custom\">加载中</div>" +
					"</div></td></tr>"
				);
				TableRow = $("#servers tr#r" + i);
				ExpandRow = $("#servers #rt" + i);
				server_status[i] = true;
			}
			TableRow = TableRow[0];
			if(error) {
				TableRow.setAttribute("data-target", "#rt" + i);
				server_status[i] = true;
			}

			// Online4
			if (result.servers[i].online4 && !result.servers[i].online6) {
				TableRow.children["online4"].children[0].children[0].className = "progress-bar progress-bar-success";
				TableRow.children["online4"].children[0].children[0].innerHTML = "<small>IPv4</small>";
			} else if (result.servers[i].online4 && result.servers[i].online6) {
				TableRow.children["online4"].children[0].children[0].className = "progress-bar progress-bar-success";
				TableRow.children["online4"].children[0].children[0].innerHTML = "<small>双栈</small>";
			} else if (!result.servers[i].online4 && result.servers[i].online6) {
			    TableRow.children["online4"].children[0].children[0].className = "progress-bar progress-bar-success";
				TableRow.children["online4"].children[0].children[0].innerHTML = "<small>IPv6</small>";
			} else {
				TableRow.children["online4"].children[0].children[0].className = "progress-bar progress-bar-danger";
				TableRow.children["online4"].children[0].children[0].innerHTML = "<small>关闭</small>";
			}

			// Online6
			//if (result.servers[i].online6) {
			//	TableRow.children["online6"].children[0].children[0].className = "progress-bar progress-bar-success";
			//	TableRow.children["online6"].children[0].children[0].innerHTML = "<small>开启</small>";
			//} else {
			//	TableRow.children["online6"].children[0].children[0].className = "progress-bar progress-bar-danger";
			//	TableRow.children["online6"].children[0].children[0].innerHTML = "<small>关闭</small>";
			//}

			// Ipstatus
			// mh361 or mh370, mourn mh370, 2014-03-08 01:20　lost from all over the world.
			if (result.servers[i].ip_status) {
				TableRow.children["ip_status"].children[0].children[0].className = "progress-bar progress-bar-success";
				TableRow.children["ip_status"].children[0].children[0].innerHTML = "<small>Normal</small>";
			} else {
				TableRow.children["ip_status"].children[0].children[0].className = "progress-bar progress-bar-danger";
				TableRow.children["ip_status"].children[0].children[0].innerHTML = "<small>Lost</small>";
			}

			// Name
			TableRow.children["name"].innerHTML = result.servers[i].name;

			// Type
			TableRow.children["type"].innerHTML = result.servers[i].type;

			//Host
			TableRow.children["host"].innerHTML = result.servers[i].host;

			// Location
			TableRow.children["location"].innerHTML = result.servers[i].location;
			if (!result.servers[i].online4 && !result.servers[i].online6) {
				if (server_status[i]) {
					TableRow.children["uptime"].innerHTML = "–";
					TableRow.children["load"].innerHTML = "–";
					TableRow.children["network"].innerHTML = "–";
					TableRow.children["traffic"].innerHTML = "–";
					TableRow.children["cpu"].children[0].children[0].className = "progress-bar progress-bar-danger";
					TableRow.children["cpu"].children[0].children[0].style.width = "100%";
					TableRow.children["cpu"].children[0].children[0].innerHTML = "<small>关闭</small>";
					TableRow.children["memory"].children[0].children[0].className = "progress-bar progress-bar-danger";
					TableRow.children["memory"].children[0].children[0].style.width = "100%";
					TableRow.children["memory"].children[0].children[0].innerHTML = "<small>关闭</small>";
					TableRow.children["hdd"].children[0].children[0].className = "progress-bar progress-bar-danger";
					TableRow.children["hdd"].children[0].children[0].style.width = "100%";
					TableRow.children["hdd"].children[0].children[0].innerHTML = "<small>关闭</small>";
					TableRow.children["ping"].children[0].children[0].className = "progress-bar progress-bar-danger";
					TableRow.children["ping"].children[0].children[0].style.width = "100%";
					TableRow.children["ping"].children[0].children[0].innerHTML = "<small>关闭</small>";
					if(ExpandRow.hasClass("in")) {
						ExpandRow.collapse("hide");
					}
					TableRow.setAttribute("data-target", "");
					server_status[i] = false;
				}
			} else {
				if (!server_status[i]) {
					TableRow.setAttribute("data-target", "#rt" + i);
					server_status[i] = true;
				}

				// Uptime
				TableRow.children["uptime"].innerHTML = getDuration(result.servers[i].uptime);

				// Load: default load_1, you can change show: load_1, load_5, load_15
				if(result.servers[i].load == -1) {
				    TableRow.children["load"].innerHTML = "–";
				} else {
				    TableRow.children["load"].innerHTML = result.servers[i].load_1.toFixed(2);
				}

				// Network
				var netstr = "";
				if(result.servers[i].network_rx < 1024)
					netstr += result.servers[i].network_rx.toFixed(0) + "B";
				else if(result.servers[i].network_rx < 1024*1024)
					netstr += (result.servers[i].network_rx/1024).toFixed(0) + "K";
				else
					netstr += (result.servers[i].network_rx/1024/1024).toFixed(1) + "M";
				netstr += " | "
				if(result.servers[i].network_tx < 1024)
					netstr += result.servers[i].network_tx.toFixed(0) + "B";
				else if(result.servers[i].network_tx < 1024*1024)
					netstr += (result.servers[i].network_tx/1024).toFixed(0) + "K";
				else
					netstr += (result.servers[i].network_tx/1024/1024).toFixed(1) + "M";
				TableRow.children["network"].innerHTML = netstr;

				//Traffic
				var trafficstr = "";
				if(result.servers[i].network_in < 1024)
					trafficstr += result.servers[i].network_in.toFixed(0) + "B";
				else if(result.servers[i].network_in < 1024*1024)
					trafficstr += (result.servers[i].network_in/1024).toFixed(0) + "K";
				else if(result.servers[i].network_in < 1024*1024*1024)
					trafficstr += (result.servers[i].network_in/1024/1024).toFixed(1) + "M";
				else if(result.servers[i].network_in < 1024*1024*1024*1024)
					trafficstr += (result.servers[i].network_in/1024/1024/1024).toFixed(2) + "G";
                else
                    trafficstr += (result.servers[i].network_in/1024/1024/1024/1024).toFixed(2) + "T";
				trafficstr += " | "
				if(result.servers[i].network_out < 1024)
					trafficstr += result.servers[i].network_out.toFixed(0) + "B";
				else if(result.servers[i].network_out < 1024*1024)
					trafficstr += (result.servers[i].network_out/1024).toFixed(0) + "K";
				else if(result.servers[i].network_out < 1024*1024*1024)
					trafficstr += (result.servers[i].network_out/1024/1024).toFixed(1) + "M";
				else if(result.servers[i].network_out < 1024*1024*1024*1024)
				    trafficstr += (result.servers[i].network_out/1024/1024/1024).toFixed(2) + "G";
				else
					trafficstr += (result.servers[i].network_out/1024/1024/1024/1024).toFixed(2) + "T";
				TableRow.children["traffic"].innerHTML = trafficstr;

				// CPU
				result.servers[i].cpu = result.servers[i].cpu.toFixed(0)
				if (result.servers[i].cpu >= 90)
					TableRow.children["cpu"].children[0].children[0].className = "progress-bar progress-bar-danger";
				else if (result.servers[i].cpu >= 80)
					TableRow.children["cpu"].children[0].children[0].className = "progress-bar progress-bar-warning";
				else
					TableRow.children["cpu"].children[0].children[0].className = "progress-bar progress-bar-success";
				TableRow.children["cpu"].children[0].children[0].style.width = result.servers[i].cpu + "%";
				TableRow.children["cpu"].children[0].children[0].innerHTML = result.servers[i].cpu + "%";

				// Memory
				var Mem = ((result.servers[i].memory_used/result.servers[i].memory_total)*100.0).toFixed(0);
				if (Mem >= 90)
					TableRow.children["memory"].children[0].children[0].className = "progress-bar progress-bar-danger";
				else if (Mem >= 80)
					TableRow.children["memory"].children[0].children[0].className = "progress-bar progress-bar-warning";
				else
					TableRow.children["memory"].children[0].children[0].className = "progress-bar progress-bar-success";
				TableRow.children["memory"].children[0].children[0].style.width = Mem + "%";
				TableRow.children["memory"].children[0].children[0].innerHTML = Mem + "%";
				ExpandRow[0].children["expand_mem"].innerHTML = "内存: " + bytesToSize(result.servers[i].memory_used*1024, 2) + "  已用 / " +bytesToSize(result.servers[i].memory_total*1024-result.servers[i].memory_used*1024, 2) + " 可用 / " + bytesToSize(result.servers[i].memory_total*1024, 2) + " 总共 ";
				// Swap
				ExpandRow[0].children["expand_swap"].innerHTML = "交换分区: " + bytesToSize(result.servers[i].swap_used*1024, 2) + " 已用 / " + bytesToSize(result.servers[i].swap_total*1024-result.servers[i].swap_used*1024, 2)+ "  可用 / "+  bytesToSize(result.servers[i].swap_total*1024, 2) + " 总共 ";

				// HDD
				var HDD = ((result.servers[i].hdd_used/result.servers[i].hdd_total)*100.0).toFixed(0);
				if (HDD >= 90)
					TableRow.children["hdd"].children[0].children[0].className = "progress-bar progress-bar-danger";
				else if (HDD >= 80)
					TableRow.children["hdd"].children[0].children[0].className = "progress-bar progress-bar-warning";
				else
					TableRow.children["hdd"].children[0].children[0].className = "progress-bar progress-bar-success";
				TableRow.children["hdd"].children[0].children[0].style.width = HDD + "%";
				TableRow.children["hdd"].children[0].children[0].innerHTML = HDD + "%";
				ExpandRow[0].children["expand_hdd"].innerHTML = "硬盘: " + bytesToSize(result.servers[i].hdd_used*1024*1024, 2) + " 已用 / " + bytesToSize(result.servers[i].hdd_total*1024*1024-result.servers[i].hdd_used*1024*1024, 2)+ "  可用 / "+ bytesToSize(result.servers[i].hdd_total*1024*1024, 2) +"  总共"

                // delay time

				// tcp, udp, process, thread count
				ExpandRow[0].children["expand_tupd"].innerHTML = "TCP/UDP/进/线: " + result.servers[i].tcp_count + " / " + result.servers[i].udp_count + " / " + result.servers[i].process_count+ " / " + result.servers[i].thread_count;
				ExpandRow[0].children["expand_ping"].innerHTML = "联通/电信/移动: " + result.servers[i].time_10010 + "ms / " + result.servers[i].time_189 + "ms / " + result.servers[i].time_10086 + "ms"

                // ping
                var PING_10010 = result.servers[i].ping_10010.toFixed(0);
                var PING_189 = result.servers[i].ping_189.toFixed(0);
                var PING_10086 = result.servers[i].ping_10086.toFixed(0);
                if (PING_10010 >= 20 || PING_189 >= 20 || PING_10086 >= 20)
                    TableRow.children["ping"].children[0].children[0].className = "progress-bar progress-bar-danger";
                else
                    TableRow.children["ping"].children[0].children[0].className = "progress-bar progress-bar-success";
	            TableRow.children["ping"].children[0].children[0].innerHTML = PING_10010 + "%💻" + PING_189 + "%💻" + PING_10086 + "%";

				// Custom
				if (result.servers[i].custom) {
					ExpandRow[0].children["expand_custom"].innerHTML = result.servers[i].custom
				} else {
					ExpandRow[0].children["expand_custom"].innerHTML = ""
				}
			}
		};

		d = new Date(result.updated*1000);
		error = 0;
	}).fail(function(update_error) {
		if (!error) {
			$("#servers > tr.accordion-toggle").each(function(i) {
				var TableRow = $("#servers tr#r" + i)[0];
				var ExpandRow = $("#servers #rt" + i);
				TableRow.children["online4"].children[0].children[0].className = "progress-bar progress-bar-error";
				TableRow.children["online4"].children[0].children[0].innerHTML = "<small>错误</small>";
				//TableRow.children["online6"].children[0].children[0].className = "progress-bar progress-bar-error";
				//TableRow.children["online6"].children[0].children[0].innerHTML = "<small>错误</small>";
				TableRow.children["ip_status"].children[0].children[0].className = "progress-bar progress-bar-error";
				TableRow.children["ip_status"].children[0].children[0].innerHTML = "<small>错误</small>";
				TableRow.children["uptime"].children[0].children[0].className = "progress-bar progress-bar-error";
				TableRow.children["uptime"].children[0].children[0].innerHTML = "<small>错误</small>";
				TableRow.children["load"].children[0].children[0].className = "progress-bar progress-bar-error";
				TableRow.children["load"].children[0].children[0].innerHTML = "<small>错误</small>";
				TableRow.children["network"].children[0].children[0].className = "progress-bar progress-bar-error";
				TableRow.children["network"].children[0].children[0].innerHTML = "<small>错误</small>";
				TableRow.children["traffic"].children[0].children[0].className = "progress-bar progress-bar-error";
				TableRow.children["traffic"].children[0].children[0].innerHTML = "<small>错误</small>";
				TableRow.children["cpu"].children[0].children[0].className = "progress-bar progress-bar-error";
				TableRow.children["cpu"].children[0].children[0].style.width = "100%";
				TableRow.children["cpu"].children[0].children[0].innerHTML = "<small>错误</small>";
				TableRow.children["memory"].children[0].children[0].className = "progress-bar progress-bar-error";
				TableRow.children["memory"].children[0].children[0].style.width = "100%";
				TableRow.children["memory"].children[0].children[0].innerHTML = "<small>错误</small>";
				TableRow.children["hdd"].children[0].children[0].className = "progress-bar progress-bar-error";
				TableRow.children["hdd"].children[0].children[0].style.width = "100%";
				TableRow.children["hdd"].children[0].children[0].innerHTML = "<small>错误</small>";
				TableRow.children["ping"].children[0].children[0].className = "progress-bar progress-bar-error";
				TableRow.children["ping"].children[0].children[0].style.width = "100%";
				TableRow.children["ping"].children[0].children[0].innerHTML = "<small>错误</small>";
				if(ExpandRow.hasClass("in")) {
					ExpandRow.collapse("hide");
				}
				TableRow.setAttribute("data-target", "");
				server_status[i] = false;
			});
		}
		error = 1;
		$("#updated").html("更新错误.");
	});
}

function updateTime() {
	if (!error)
		$("#updated").html("最后更新: " + timeSince(d));
}

uptime();
updateTime();
setInterval(uptime, 2000);
setInterval(updateTime, 2000);


// styleswitcher.js
function setActiveStyleSheet(title) {
	var i, a, main;
	for(i=0; (a = document.getElementsByTagName("link")[i]); i++) {
		if(a.getAttribute("rel").indexOf("style") != -1 && a.getAttribute("title")) {
			a.disabled = true;
			if(a.getAttribute("title") == title) a.disabled = false;
		}
	}
}

function getActiveStyleSheet() {
	var i, a;
	for(i=0; (a = document.getElementsByTagName("link")[i]); i++) {
		if(a.getAttribute("rel").indexOf("style") != -1 && a.getAttribute("title") && !a.disabled)
			return a.getAttribute("title");
	}
	return null;
}

function getPreferredStyleSheet() {
	var i, a;
	for(i=0; (a = document.getElementsByTagName("link")[i]); i++) {
		if(a.getAttribute("rel").indexOf("style") != -1	&& a.getAttribute("rel").indexOf("alt") == -1 && a.getAttribute("title"))
			return a.getAttribute("title");
	}
return null;
}

function createCookie(name,value,days) {
	if (days) {
		var date = new Date();
		date.setTime(date.getTime()+(days*24*60*60*1000));
		var expires = "; expires="+date.toGMTString();
	}
	else expires = "";
	document.cookie = name+"="+value+expires+"; path=/";
}

function readCookie(name) {
	var nameEQ = name + "=";
	var ca = document.cookie.split(';');
	for(var i=0;i < ca.length;i++) {
		var c = ca[i];
		while (c.charAt(0)==' ')
			c = c.substring(1,c.length);
		if (c.indexOf(nameEQ) == 0)
			return c.substring(nameEQ.length,c.length);
	}
	return null;
}

window.onload = function(e) {
	var cookie = readCookie("style");
	var title = cookie ? cookie : getPreferredStyleSheet();
	setActiveStyleSheet(title);
}

window.onunload = function(e) {
	var title = getActiveStyleSheet();
	createCookie("style", title, 365);
}

var cookie = readCookie("style");
var title = cookie ? cookie : getPreferredStyleSheet();
setActiveStyleSheet(title);

// Excel 报表导出，参考https://www.cnblogs.com/zhenggaowei/p/11732170.html
var fileName = '';


function json2Excel() {
  var wopts = {
	bookType: 'xlsx',
	bookSST: false,
	type: 'binary'
  };
  var workBook = {
	SheetNames: ['Sheet1'],
	Sheets: {},
	Props: {}
  };
   /** 
   		TODO: 
   		多余信息现在隐藏不导出，后期再优化 
   		Author:CHN-STUDENT <chn-student@outlook.com> 2020.12.01
   */
  
   var elements1 =  document.getElementsByClassName("expandRow even");
   var elements2 = document.getElementsByClassName("expandRow odd");
   console.log(elements2)
   Array.prototype.forEach.call(elements1, function (element) {
		element.style.display = 'none';	
   });
   Array.prototype.forEach.call(elements2, function (element) {
		element.style.display = 'none';	
   });
   workBook.Sheets['Sheet1'] =  XLSX.utils.table_to_sheet(document.getElementById('info'),{display:true});
    
	//3、XLSX.write() 开始编写Excel表格
	//4、changeData() 将数据处理成需要输出的格式
	saveAs(new Blob([changeData(XLSX.write(workBook, wopts))], { type: 'application/octet-stream' }))
	Array.prototype.forEach.call(elements1, function (element) {
		element.style = '';	
   });
   Array.prototype.forEach.call(elements2, function (element) {
		element.style = '';	
   });
}
function changeData(s) {
  //如果存在ArrayBuffer对象(es6) 最好采用该对象
  if (typeof ArrayBuffer !== 'undefined') {

	//1、创建一个字节长度为s.length的内存区域
	var buf = new ArrayBuffer(s.length);

	//2、创建一个指向buf的Unit8视图，开始于字节0，直到缓冲区的末尾
	var view = new Uint8Array(buf);

	//3、返回指定位置的字符的Unicode编码
	for (var i = 0; i != s.length; ++i) view[i] = s.charCodeAt(i) & 0xFF;
	return buf;

  } else {
	var buf = new Array(s.length);
	for (var i = 0; i != s.length; ++i) buf[i] = s.charCodeAt(i) & 0xFF;
	return buf;
  }
}
function saveAs(obj, fileName) {//当然可以自定义简单的下载文件实现方式
	var tmpa = document.createElement("a");
	tmpa.download = fileName ? fileName + '.xlsx' : new Date().getTime() + '.xlsx';
	tmpa.href = URL.createObjectURL(obj); //绑定a标签
	tmpa.click(); //模拟点击实现下载

	setTimeout(function () { //延时释放
	URL.revokeObjectURL(obj); //用URL.revokeObjectURL()来释放这个object URL
	}, 100);
}

