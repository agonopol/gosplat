package gosplat

var HTML = `<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01//EN">
<html>
<head>
	<title>{{.Title}}</title>
	<link rel="stylesheet" href="http://yui.yahooapis.com/2.9.0/build/reset/reset-min.css" type="text/css" media="screen" charset="utf-8">
	<style>
	/* ================ */
	/* = The 1Kb Grid = */     /* 12 columns, 80 pixels each, with 20 pixel gutter */
	/* ================ */

	.grid_1 { width:80px; }
	.grid_2 { width:180px; }
	.grid_3 { width:280px; }
	.grid_4 { width:380px; }
	.grid_5 { width:480px; }
	.grid_6 { width:580px; }
	.grid_7 { width:680px; }
	.grid_8 { width:780px; }
	.grid_9 { width:880px; }
	.grid_10 { width:980px; }
	.grid_11 { width:1080px; }
	.grid_12 { width:1180px; }

	.column {
		margin: 0 10px;
		overflow: hidden;
		float: left;
		display: inline;
	}
	.row {
		width: 1200px;
		margin: 0 auto;
		overflow: hidden;
	}
	.row .row {
		margin: 0 -10px;
		width: auto;
		display: inline-block;
	}

	body {
		background: url(data:image/gif;base64,R0lGODlhJAAkAIAAADMzMwAAACH5BAAAAAAALAAAAAAkACQAAAK7RH5mocn/VkO0rQiR1Zj2zXHT5DgK2ZkgdG2XJ6ZSZlbA/aLRSt6eXwvqKgncDKaKEYEanXO4a9JEOWemNT0Vj0gXssQMXzG9q4SZqlmrMG9pvAoly1gtOzZ7KeP4/hLOEpj3QzhIw5NG5pfzMSWUyOhDh4eip4j15tYUWQh52FLXN1m09dam+HUmpVcWRGXjGEVqGRW6xLr2l7oFRDsHSArrOpL1iaO7J/qlDEZ5+Cw5llz5Z1jthw1SAAA7);
		font-family: helvetica;
		font-size: 14px;
	}

	h1 {
		font-size: 2.0em;
		border-bottom: solid #CCC 4px;
		margin-bottom: 1.5em;
	}

	h3 {
		text-align: center;
	}

	h1, h3 {
		font-weight: bold;
		margin-left: 0.5em;
	}

	#content {
		background-color: #F1F1F1;
		padding-left: 1em;
		padding-bottom: 10em;
	}

	.graphrow {
		padding-bottom: 3em;
	}

	.row {
		background-color: #FFF;
	}
	
	</style>
	<script type="text/javascript" src="https://www.google.com/jsapi"></script>
	<script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.6.2/jquery.min.js"></script>
	<script type="text/javascript" src="http://dygraphs.com/dygraph-combined.js"></script>
	<script type="text/javascript" src="http://dygraphs.com/dygraph-options.js"></script>
	<script type="text/javascript" charset="utf-8">
	google.load("visualization", "1", {packages:["corechart", "table", "treemap"]});
	</script>
	
	<script language="javascript" type="text/javascript">
	function popitup(url, params) {
		newwindow=window.open(url,'name', params);
		if (window.focus) {newwindow.focus()}
		return false;
	}
	</script>
	
	<!--Barchart-->
	<script type="text/javascript" charset="utf-8">
	var Barchart = {};

	Barchart.plot = function(container, data, opts) {
		var chartdata = new google.visualization.DataTable();
		var i = 0;

		for (key in data[0]) {
			var t = typeof data[0][key];
			chartdata.addColumn(t, key);
		}

		chartdata.addRows(data.length);

		for (var i=0; i < data.length; i++) {
			var j = 0;
			for (key in data[i]) {
				chartdata.setValue(i, j, data[i][key]);
				j++;
			}
		};

		var chart = new google.visualization.BarChart(document.getElementById(container));
		chart.draw(chartdata, jQuery.extend({width: '1200px', height: '500px'}, opts));
	}
	</script>
	
	<!--Candlestick-->
	<script type="text/javascript" charset="utf-8">
	var Candlestick = {};

	Candlestick.plot = function(container, data, opts) {
		var dataTable = google.visualization.arrayToDataTable(data, true);
		var chart = new google.visualization.CandlestickChart(document.getElementById(container));
		chart.draw(dataTable, jQuery.extend({legend: 'none'}, opts));
	}
	</script>
	
	<!--Columnchart-->
	<script type="text/javascript" charset="utf-8">
	var Columnchart = {};

	Columnchart.plot = function(container, data, opts) {
		var chartdata = new google.visualization.DataTable();
		var i = 0;

		for (key in data[0]) {
			var t = typeof data[0][key];
			chartdata.addColumn(t, key);
		}

		chartdata.addRows(data.length);

		for (var i=0; i < data.length; i++) {
			var j = 0;
			for (key in data[i]) {
				chartdata.setValue(i, j, data[i][key]);
				j++;
			}
		};

		var chart = new google.visualization.ColumnChart(document.getElementById(container)); chart.draw(chartdata, jQuery.extend({width:
			'1200px', height: '500px'}, opts)); }
	</script>
	  
	<!--Combochart-->
	<script type="text/javascript" charset="utf-8">
	var Combochart = {};

	Combochart.plot = function(container, data, opts) {
		var chartdata = new google.visualization.DataTable();
		var i = 0;

		for (key in data[0]) {
			var t = typeof data[0][key];
			chartdata.addColumn(t, key);
		}

		chartdata.addRows(data.length);

		for (var i=0; i < data.length; i++) {
			var j = 0;
			for (key in data[i]) {
				chartdata.setValue(i, j, data[i][key]);
				j++;
			}
		};

		var chart = new google.visualization.ComboChart(document.getElementById(container));
		chart.draw(chartdata, jQuery.extend({width: '1200px', height: '500px'}, opts));
	}
	</script>
	 
	<!--Linechart-->
	<script type="text/javascript" charset="utf-8">
	var Linechart = {};

	Linechart.plot = function(container, data, opts) {
		var chartdata = new google.visualization.DataTable();
		var i = 0;

		for (key in data[0]) {
			var t = typeof data[0][key];
			chartdata.addColumn(t, key);
		}

		chartdata.addRows(data.length);

		for (var i=0; i < data.length; i++) {
			var j = 0;
			for (key in data[i]) {
				chartdata.setValue(i, j, data[i][key]);
				j++;
			}
		};

		var chart = new google.visualization.LineChart(document.getElementById(container));
		chart.draw(chartdata, opts);
	}
	</script>
	 
	<!--Piechart-->
	<script type="text/javascript" charset="utf-8">
	var Piechart = {};

	Piechart.plot = function(container, data, opts) {
		var chartdata = new google.visualization.DataTable();
		var i = 0;

		for (key in data[0]) {
			var t = typeof data[0][key];
			chartdata.addColumn(t, key);
		}

		chartdata.addRows(data.length);

		for (var i=0; i < data.length; i++) {
			var j = 0;
			for (key in data[i]) {
				chartdata.setValue(i, j, data[i][key]);
				j++;
			}
		};

		var chart = new google.visualization.PieChart(document.getElementById(container));
		chart.draw(chartdata, jQuery.extend({width: '1200px', height: '500px'}, opts));
	}
	</script>
	 
	<!--Scatter--> 
	<script type="text/javascript" charset="utf-8"> var Scatter = {};

	Scatter.plot = function(container, data, opts) {
		var chartdata = new google.visualization.DataTable();
		var plotvals = {};
		var cols = []
		var types = []
		var length = 0
		for (var i=0; i < data.length; i++) {
			d = data[i]
			for (key in d) {
				if (key != 'x') {
					if (cols.indexOf(key) == -1) {
						cols.push(key)
						types.push(typeof d[key])
					}
				}
				if (!(d['x'] in plotvals)) {
					length++
					plotvals[d['x']] = {}
				} 
				plotvals[d['x']][key] = d[key]        
			}
		}

		var xtype = typeof data[0]['x']
		chartdata.addColumn(xtype, 'x')
		for (key in cols) {
			chartdata.addColumn(types[key], cols[key]);
		}

		chartdata.addRows(length);

		var i = 0;
		for (key in plotvals) {
			chartdata.setValue(i, 0, plotvals[key]['x']);

			for (col in plotvals[key]) {
				chartdata.setValue(i, cols.indexOf(col) + 1, plotvals[key][col]);
			}
			i++
		}
		var chart = new google.visualization.ScatterChart(document.getElementById(container));
		chart.draw(chartdata, jQuery.extend({width: '700px', height: '700px'}, opts));
	}
	</script>
	 
	<!--Table-->
	<script type="text/javascript" charset="utf-8">
	var Table = {};

	Table.plot = function(container, data, opts) {
		var chartdata = new google.visualization.DataTable();
		var i = 0;

		for (key in data[0]) {
			var t = typeof data[0][key];
			chartdata.addColumn(t, key);
		}

		chartdata.addRows(data.length);

		for (var i=0; i < data.length; i++) {
			var j = 0;
			for (key in data[i]) {
				chartdata.setValue(i, j, data[i][key]);
				j++;
			}
		};

		var chart = new google.visualization.Table(document.getElementById(container));
		chart.draw(chartdata, jQuery.extend({width: '1200'}, opts));
	}
	</script>
	 
	<!--Timeseries-->
	<script type="text/javascript" charset="utf-8">
	var Timeseries = {};

	Timeseries.plot = function(container, data, opts) {
		var plotvals = {};
		var cols = []
		var length = 0
		for (var i=0; i < data.length; i++) {
			d = data[i]
			for (key in d) {
				if (key != 'x') {
					if (cols.indexOf(key) == -1) {
						cols.push(key)
					}
				}
				if (!(d['x'] in plotvals)) {
					length++
					plotvals[d['x']] = {}
				} 
				plotvals[d['x']][key] = d[key]        
			}
		}

		var csv = 'x,' + cols.join(',') + "\n";

		for (key in plotvals) {
			var vals = []
			for (col_ind in cols) {
				col = cols[col_ind]
				if (col in plotvals[key]) {
					vals.push(plotvals[key][col])
				} else {
					vals.push('')
				}
			}
			csv += plotvals[key]['x'] + ',' + vals.join(',') + '\n'
		}

		g = new Dygraph(document.getElementById(container), csv, opts);
	}
	</script>
	 
	<!--Treemap-->
	<script type="text/javascript" charset="utf-8">
	var Treemap = {};

	Treemap.plot = function(container, data, opts) {
		var chartdata = new google.visualization.DataTable();
		var i = 0;

		for (key in data[0]) {
			if (data[0][key] == null) {
				var t = "string";
			} else {
				var t = typeof data[0][key];
			}
			chartdata.addColumn(t, key);
		}

		var rows = [];

		for (var i=0; i < data.length; i++) {
			var row = []
			var j = 0;
			for (key in data[i]) {
				row[j] = data[i][key];
				j++;
			}
			console.log(row);
			rows[i] = row;
		}

		chartdata.addRows(rows);

		var chart = new google.visualization.TreeMap(document.getElementById(container));
		chart.draw(chartdata, jQuery.extend({width: '1200px', height: '500px'}, opts));
	}
	</script>
</head>
<body>
<div id="content">
  <div class="row">
    <div id="header" class="column grid_12"><h1>{{.Title}}</h1></div>
  </div>
	{{with .Rows}}{{range .}}
  	<div class="row graphrow">{{with .Visualizations}}{{range .}}
  		<div class="column {{.ColWidth}}">
			<h3>{{.Title}}</h3>
		<div id="{{.Id}}" style="width:100%; height:{{.Height}};">
	  </div>
	</div>{{end}}{{end}}
  </div>{{end}}{{end}}
</div>
</body>
  {{with .Rows}}{{range .}}
  {{with .Visualizations}}{{range .}}
  <script>eval("{{.Class}}").plot("{{.Id}}", {{.Data}}, {{.Options}})</script>
  {{end}}{{end}}
  {{end}}{{end}}  
</html>`
